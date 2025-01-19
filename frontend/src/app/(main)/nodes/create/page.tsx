'use client'

import { Button } from '@/components/ui/Button'
import { Plus, Minus } from 'lucide-react'
import { useNodeForm } from '@/application/hooks/useNodeForm'

export default function CreateNodePage() {
  const {
    formData,
    currentStep,
    handleInputChange,
    handleRewardChange,
    addReward,
    removeReward,
    nextStep,
    prevStep
  } = useNodeForm()

  return (
    <div className="max-w-4xl mx-auto py-8 px-4">
      <h1 className="text-3xl font-bold mb-8">Crear Nuevo Nodo</h1>
      
      {/* Progress Steps */}
      <div className="flex justify-between mb-8">
        <div className={`step ${currentStep >= 1 ? 'active' : ''}`}>Información Básica</div>
        <div className={`step ${currentStep >= 2 ? 'active' : ''}`}>Detalles del Proyecto</div>
        <div className={`step ${currentStep >= 3 ? 'active' : ''}`}>Recompensas</div>
      </div>

      {/* Step 1: Basic Information */}
      {currentStep === 1 && (
        <div className="space-y-6">
          <div>
            <label htmlFor="title" className="block mb-2">Título del Proyecto</label>
            <input
              id="title"
              type="text"
              value={formData.title}
              onChange={(e) => handleInputChange('title', e.target.value)}
              className="w-full p-2 border rounded"
              aria-label="Título del proyecto"
              placeholder="Ingrese el título del proyecto"
            />
          </div>
          
          <div>
            <label htmlFor="category" className="block mb-2">Categoría</label>
            <select
              id="category"
              value={formData.category}
              onChange={(e) => handleInputChange('category', e.target.value)}
              className="w-full p-2 border rounded"
              aria-label="Categoría del proyecto"
            >
              <option value="">Seleccionar categoría</option>
              <option value="social">Social</option>
              <option value="ambiental">Ambiental</option>
              <option value="animal">Animal</option>
            </select>
          </div>

          <div>
            <label htmlFor="location" className="block mb-2">Ubicación</label>
            <input
              id="location"
              type="text"
              value={formData.location}
              onChange={(e) => handleInputChange('location', e.target.value)}
              className="w-full p-2 border rounded"
              aria-label="Ubicación del proyecto"
              placeholder="Ingrese la ubicación"
            />
          </div>

          <div className="flex justify-end">
            <Button onClick={nextStep}>Siguiente</Button>
          </div>
        </div>
      )}

      {/* Step 2: Project Details */}
      {currentStep === 2 && (
        <div className="space-y-6">
          <div>
            <label htmlFor="goal" className="block mb-2">Meta de Financiamiento</label>
            <input
              id="goal"
              type="number"
              value={formData.goal}
              onChange={(e) => handleInputChange('goal', e.target.value)}
              className="w-full p-2 border rounded"
              aria-label="Meta de financiamiento"
              placeholder="Ingrese la meta"
            />
          </div>

          <div>
            <label htmlFor="duration" className="block mb-2">Duración de la Campaña</label>
            <input
              id="duration"
              type="number"
              value={formData.duration}
              onChange={(e) => handleInputChange('duration', e.target.value)}
              className="w-full p-2 border rounded"
              aria-label="Duración de la campaña"
              placeholder="Ingrese la duración en días"
            />
          </div>

          <div>
            <label htmlFor="description" className="block mb-2">Descripción del Proyecto</label>
            <textarea
              id="description"
              value={formData.description}
              onChange={(e) => handleInputChange('description', e.target.value)}
              className="w-full p-2 border rounded h-32"
              aria-label="Descripción del proyecto"
              placeholder="Describa su proyecto"
            />
          </div>

          <div className="flex justify-between">
            <Button onClick={prevStep} variant="ghost">Anterior</Button>
            <Button onClick={nextStep}>Siguiente</Button>
          </div>
        </div>
      )}

      {/* Step 3: Rewards */}
      {currentStep === 3 && (
        <div className="space-y-6">
          {formData.rewards.map((reward, index) => (
            <div key={index} className="p-4 border rounded space-y-4">
              <div className="flex justify-between items-center">
                <h3 className="text-lg font-semibold">Recompensa {index + 1}</h3>
                {index > 0 && (
                  <Button onClick={() => removeReward(index)} variant="ghost" size="icon">
                    <Minus className="h-4 w-4" />
                  </Button>
                )}
              </div>

              <div>
                <label htmlFor={`reward-title-${index}`} className="block mb-2">Título</label>
                <input
                  id={`reward-title-${index}`}
                  type="text"
                  value={reward.title}
                  onChange={(e) => handleRewardChange(index, 'title', e.target.value)}
                  className="w-full p-2 border rounded"
                  aria-label={`Título de la recompensa ${index + 1}`}
                  placeholder="Ingrese el título de la recompensa"
                />
              </div>

              <div>
                <label htmlFor={`reward-amount-${index}`} className="block mb-2">Monto</label>
                <input
                  id={`reward-amount-${index}`}
                  type="number"
                  value={reward.amount}
                  onChange={(e) => handleRewardChange(index, 'amount', e.target.value)}
                  className="w-full p-2 border rounded"
                  aria-label={`Monto de la recompensa ${index + 1}`}
                  placeholder="Ingrese el monto"
                />
              </div>

              <div>
                <label htmlFor={`reward-description-${index}`} className="block mb-2">Descripción</label>
                <textarea
                  id={`reward-description-${index}`}
                  value={reward.description}
                  onChange={(e) => handleRewardChange(index, 'description', e.target.value)}
                  className="w-full p-2 border rounded h-24"
                  aria-label={`Descripción de la recompensa ${index + 1}`}
                  placeholder="Describa la recompensa"
                />
              </div>
            </div>
          ))}

          <Button onClick={addReward} variant="ghost" className="w-full">
            <Plus className="h-4 w-4 mr-2" />
            Agregar Recompensa
          </Button>

          <div className="flex justify-between mt-8">
            <Button onClick={prevStep} variant="ghost">Anterior</Button>
            <Button onClick={() => console.log(formData)}>Crear Nodo</Button>
          </div>
        </div>
      )}
    </div>
  )
}
