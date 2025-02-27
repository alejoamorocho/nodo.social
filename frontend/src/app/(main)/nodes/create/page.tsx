'use client'

import { useState } from 'react';
import { Button } from '@/components/ui/Button';
import { Plus, Minus } from 'lucide-react';
import { db, storage, auth } from '../../../../firebase'; 
import { collection, addDoc } from 'firebase/firestore';
import { ref, uploadBytes, getDownloadURL } from 'firebase/storage';
import Modal from '@/components/ui/Modal';

interface Reward {
  title: string;
  amount: string;
  description: string;
}

interface FormData {
  title: string;
  category: string;
  location: string;
  goal: string;
  duration: string;
  description: string;
  rewards: Reward[];
  image: File | null;
}

export default function CreateNodePage() {
  const [formData, setFormData] = useState<FormData>({
    title: '',
    category: '',
    location: '',
    goal: '',
    duration: '',
    description: '',
    rewards: [],
    image: null,
  });
  const [currentStep, setCurrentStep] = useState(1);
  const [error, setError] = useState<string | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [modalMessage, setModalMessage] = useState('');
  

  const handleInputChange = (field: keyof FormData, value: string | number | File | null) => {
    setFormData({ ...formData, [field]: value });
  };

  const handleRewardChange = (index: number, field: keyof Reward, value: string | number) => {
    const newRewards = [...formData.rewards];
    newRewards[index] = { ...newRewards[index], [field]: value };
    setFormData({ ...formData, rewards: newRewards });
  };

  const addReward = () => {
    setFormData({ ...formData, rewards: [...formData.rewards, { title: '', amount: '', description: '' }] });
  };

  const removeReward = (index: number) => {
    const newRewards = formData.rewards.filter((_, i) => i !== index);
    setFormData({ ...formData, rewards: newRewards });
  };

  const nextStep = () => {
    if (currentStep === 1 && (!formData.title || !formData.category || !formData.location)) {
      setError('Por favor, complete todos los campos obligatorios.');
      return;
    }
    if (currentStep === 2 && (!formData.goal || !formData.duration || !formData.description)) {
      setError('Por favor, complete todos los campos obligatorios.');
      return;
    }
    setError(null);
    setCurrentStep(currentStep + 1);
  };

  const prevStep = () => {
    setCurrentStep(currentStep - 1);
  };

  const handleCreateNode = async () => {
    try {
      const user = auth.currentUser;
      if (!user) {
        setModalMessage('You must be logged in to create a node.');
        setIsModalOpen(true);
        return;
      }
  
      // Subir la imagen a Firebase Storage (si existe)
      let imageUrl = '';
      if (formData.image) {
        const storageRef = ref(storage, `nodes/${formData.image.name}`);
        await uploadBytes(storageRef, formData.image);
        imageUrl = await getDownloadURL(storageRef);
      }
  
      // Guardar el nodo en Firestore
      await addDoc(collection(db, 'nodes'), {
        title: formData.title,
        category: formData.category,
        location: formData.location,
        goal: formData.goal,
        duration: formData.duration,
        description: formData.description,
        rewards: formData.rewards,
        imageUrl,
        createdAt: new Date(),
        followers: [],
        createdBy: user.uid,
      });
  
      setModalMessage('Node creado con éxito');
      setIsModalOpen(true);
    } catch (error) {
      console.error('Error:', error);
      setModalMessage('Error al crear el nodo');
      setIsModalOpen(true);
    }
  };

  const closeModal = () => {
    setIsModalOpen(false);
    if (modalMessage === 'Node creado con éxito') {
      // Redirigir a la página de nodos
      window.location.href = '/nodes';
    }
  };

  return (
    <div className="max-w-4xl mx-auto py-12 px-6">
      {/* Título y descripción */}
      <h1 className="text-4xl font-bold mb-4 text-primary">Crear Nuevo Nodo</h1>
      <p className="text-lg mb-8 text-foreground/80">
        Complete los siguientes pasos para crear un nuevo nodo.
      </p>
    
      {/* Progress Steps */}
      <div className="flex justify-between mb-8 relative">
        {[1, 2, 3].map((step) => (
          <div
            key={step}
            className={`flex-1 text-center relative ${
              currentStep >= step ? "text-primary" : "text-foreground/50"
            }`}
          >
            {/* Círculo del paso */}
            <div
              className={`w-20 h-20 mx-auto mb-2 rounded-full flex items-center justify-center border-2 ${
                currentStep >= step
                  ? "border-primary bg-primary text-background"
                  : "border-current-line bg-background text-foreground/50"
              } transition-all duration-300 `}
            >
              {step}
            </div>

            {/* Texto del paso */}
            <span className="block text-sm font-medium">
              {step === 1 && "Información Básica"}
              {step === 2 && "Detalles del Proyecto"}
              {step === 3 && "Recompensas"}
            </span>

            {/* Línea que conecta los pasos */}
            {step < 3 && (
              <div
                className={`absolute top-6 left-1/2 w-64 h-0.5 ${
                  currentStep > step ? "bg-primary" : "bg-current-line"
                } transition-all duration-300 `}
              />
            )}
          </div>
        ))}
      </div>
    
      {/* Mensaje de error */}
      {error && (
        <div className="mb-6 p-4 bg-error/10 text-error rounded-lg">{error}</div>
      )}
    
      {/* Step 1: Basic Information */}
      {currentStep === 1 && (
        <div className="space-y-6">
          <div>
            <label htmlFor="title" className="block mb-2 text-foreground">
              Título del Proyecto <span className="text-error">*</span>
            </label>
            <input
              id="title"
              type="text"
              value={formData.title}
              onChange={(e) => handleInputChange("title", e.target.value)}
              className="input w-full p-3"
              aria-label="Título del proyecto"
              placeholder="Ingrese el título del proyecto"
            />
          </div>
    
          <div>
            <label htmlFor="category" className="block mb-2 text-foreground">
              Categoría <span className="text-error">*</span>
            </label>
            <select
              id="category"
              value={formData.category}
              onChange={(e) => handleInputChange("category", e.target.value)}
              className="input w-full p-3"
              aria-label="Categoría del proyecto"
            >
              <option value="">Seleccionar categoría</option>
              <option value="social">Social</option>
              <option value="ambiental">Ambiental</option>
              <option value="animal">Animal</option>
            </select>
          </div>
    
          <div>
            <label htmlFor="location" className="block mb-2 text-foreground">
              Ubicación <span className="text-error">*</span>
            </label>
            <input
              id="location"
              type="text"
              value={formData.location}
              onChange={(e) => handleInputChange("location", e.target.value)}
              className="input w-full p-3"
              aria-label="Ubicación del proyecto"
              placeholder="Ingrese la ubicación"
            />
          </div>

          <div>
            <label htmlFor="image" className="block mb-2 text-foreground">
              Imagen del Proyecto
            </label>
            <input
              id="image"
              type="file"
              onChange={(e) => handleInputChange("image", e.target.files ? e.target.files[0] : null)}
              className="input w-full p-3"
              aria-label="Imagen del proyecto"
            />
          </div>
    
          <div className="flex justify-end">
            <Button onClick={nextStep} className="btn btn-primary">
              Siguiente
            </Button>
          </div>
        </div>
      )}
    
      {/* Step 2: Project Details */}
      {currentStep === 2 && (
        <div className="space-y-6">
          <div>
            <label htmlFor="goal" className="block mb-2 text-foreground">
              Meta de Financiamiento <span className="text-error">*</span>
            </label>
            <input
              id="goal"
              type="number"
              value={formData.goal}
              onChange={(e) => handleInputChange("goal", e.target.value)}
              className="input w-full p-3"
              aria-label="Meta de financiamiento"
              placeholder="Ingrese la meta"
            />
          </div>
    
          <div>
            <label htmlFor="duration" className="block mb-2 text-foreground">
              Duración de la Campaña <span className="text-error">*</span>
            </label>
            <input
              id="duration"
              type="number"
              value={formData.duration}
              onChange={(e) => handleInputChange("duration", e.target.value)}
              className="input w-full p-3"
              aria-label="Duración de la campaña"
              placeholder="Ingrese la duración en días"
            />
          </div>
    
          <div>
            <label htmlFor="description" className="block mb-2 text-foreground">
              Descripción del Proyecto <span className="text-error">*</span>
            </label>
            <textarea
              id="description"
              value={formData.description}
              onChange={(e) => handleInputChange("description", e.target.value)}
              className="input w-full p-3 h-32"
              aria-label="Descripción del proyecto"
              placeholder="Describa su proyecto"
            />
          </div>
    
          <div className="flex justify-between">
            <Button onClick={prevStep} className="btn btn-ghost">
              Anterior
            </Button>
            <Button onClick={nextStep} className="btn btn-primary">
              Siguiente
            </Button>
          </div>
        </div>
      )}
    
      {/* Step 3: Rewards */}
      {currentStep === 3 && (
        <div className="space-y-6">
          {formData.rewards.map((reward, index) => (
            <div key={index} className="p-6 border rounded-lg space-y-4 bg-current-line/10">
              <div className="flex justify-between items-center">
                <h3 className="text-lg font-semibold text-foreground">
                  Recompensa {index + 1}
                </h3>
                {index > 0 && (
                  <Button
                    onClick={() => removeReward(index)}
                    className="btn btn-ghost text-error hover:bg-error/10"
                  >
                    <Minus className="h-4 w-4" />
                  </Button>
                )}
              </div>
    
              <div>
                <label
                  htmlFor={`reward-title-${index}`}
                  className="block mb-2 text-foreground"
                >
                  Título
                </label>
                <input
                  id={`reward-title-${index}`}
                  type="text"
                  value={reward.title}
                  onChange={(e) => handleRewardChange(index, "title", e.target.value)}
                  className="input w-full p-3"
                  aria-label={`Título de la recompensa ${index + 1}`}
                  placeholder="Ingrese el título de la recompensa"
                />
              </div>
    
              <div>
                <label
                  htmlFor={`reward-amount-${index}`}
                  className="block mb-2 text-foreground"
                >
                  Monto
                </label>
                <input
                  id={`reward-amount-${index}`}
                  type="number"
                  value={reward.amount}
                  onChange={(e) => handleRewardChange(index, "amount", e.target.value)}
                  className="input w-full p-3"
                  aria-label={`Monto de la recompensa ${index + 1}`}
                  placeholder="Ingrese el monto"
                />
              </div>
    
              <div>
                <label
                  htmlFor={`reward-description-${index}`}
                  className="block mb-2 text-foreground"
                >
                  Descripción
                </label>
                <textarea
                  id={`reward-description-${index}`}
                  value={reward.description}
                  onChange={(e) =>
                    handleRewardChange(index, "description", e.target.value)
                  }
                  className="input w-full p-3 h-24"
                  aria-label={`Descripción de la recompensa ${index + 1}`}
                  placeholder="Describa la recompensa"
                />
              </div>
            </div>
          ))}
    
          <Button
            onClick={addReward}
            className="btn btn-ghost w-full text-primary hover:bg-primary/10"
          >
            <Plus className="h-4 w-4 mr-2" />
            Agregar Recompensa
          </Button>
    
          <div className="flex justify-between mt-8">
            <Button onClick={prevStep} className="btn btn-ghost">
              Anterior
            </Button>
            <Button onClick={handleCreateNode} className="btn btn-primary">
              Crear Nodo
            </Button>
          </div>
        </div>
      )}
      <Modal isOpen={isModalOpen} onClose={closeModal} title="Resultado">
        <p>{modalMessage}</p>
      </Modal>
    </div>
  )
}
