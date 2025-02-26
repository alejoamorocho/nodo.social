'use client';

import { Card } from '@/components/ui/Card';
import { Button } from '@/components/ui/Button';
import { useState } from 'react';
import { createUserWithEmailAndPassword, updateProfile } from 'firebase/auth';
import { auth, db } from '../../../firebase';
import { useRouter } from 'next/navigation';
import { FaEye, FaEyeSlash } from 'react-icons/fa';
import { getStorage, ref, uploadBytes, getDownloadURL } from 'firebase/storage';
import { doc, setDoc } from 'firebase/firestore';

export default function RegisterPage() {
  const [credentials, setCredentials] = useState({
    name: '',
    email: '',
    password: '',
  });
  const [nameError, setNameError] = useState('');
  const [emailError, setEmailError] = useState('');
  const [passwordError, setPasswordError] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  const [avatar, setAvatar] = useState<File | null>(null);
  const { push } = useRouter();

  const changeUser = (e: React.ChangeEvent<HTMLInputElement>) => {
    setCredentials({
      ...credentials,
      [e.target.name]: e.target.value,
    });
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      setAvatar(e.target.files[0]);
    }
  };

  const validateEmail = (email: string) => {
    const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return re.test(String(email).toLowerCase());
  };

  const validatePassword = (password: string) => {
    return password.length >= 6;
  };

  const registerUser = async (e: React.FormEvent) => {
    e.preventDefault();
    const { name, email, password } = credentials;

    // Reset errors
    setNameError('');
    setEmailError('');
    setPasswordError('');

    // Validar campos vacíos
    if (!name) {
      setNameError('Por favor, ingresa tu nombre.');
    }
    if (!email) {
      setEmailError('Por favor, ingresa tu email.');
    }
    if (!password) {
      setPasswordError('Por favor, ingresa tu contraseña.');
    }

    if (!name || !email || !password) {
      return;
    }

    // Validar email
    if (!validateEmail(email)) {
      setEmailError('Por favor, ingresa un email válido.');
      return;
    }

    // Validar contraseña
    if (!validatePassword(password)) {
      setPasswordError('La contraseña debe tener al menos 6 caracteres.');
      return;
    }

    try {
      const userCredential = await createUserWithEmailAndPassword(auth, email, password);
      const user = userCredential.user;

      let avatarURL = '';
      if (avatar) {
        const storage = getStorage();
        const storageRef = ref(storage, `avatars/${user.uid}`);
        await uploadBytes(storageRef, avatar);
        avatarURL = await getDownloadURL(storageRef);
      }

      await updateProfile(user, {
        displayName: name,
        photoURL: avatarURL,
      });

      // Guardar datos del usuario en Firestore
      const joinDate = new Date().toISOString().split('T')[0]; // Formato YYYY-MM-DD
      await setDoc(doc(db, 'users', user.uid), {
        name: name,
        email: email,
        avatar: avatarURL,
        coverImage: '',
        bio: '',
        location: '',
        joinDate: joinDate,
        links: {
          website: '',
          twitter: '',
          instagram: ''
        },
        stats: {
          nodesCreated: 0,
          nodesSupported: 0,
          totalImpact: 0,
          followers: 0
        }
      });

      push('/');
    } catch (error: unknown) {
      if (error instanceof Error) {
        setEmailError(`Hubo un problema al crear tu cuenta: ${error.message}`);
        console.log(error.message);
      } else {
        console.log('Error desconocido', error);
      }
    }
  };

  return (
    <div className="min-h-screen flex flex-col md:flex-row">
      <div className="hidden md:flex flex-1 items-center justify-center p-8 md:p-16">
        <div className="text-center">
          <h1 className="text-5xl font-bold mb-4 bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
            Bienvenido a NODO.SOCIAL
          </h1>
          <h2 className="text-2xl font-semibold mb-32 mt-24">Conectando personas y emprendedores con causas sociales, ambientales y animales a través de nodos temáticos.</h2>
          <p className="text-lg mb-8">Crea tu cuenta para empezar a participar.</p>
        </div>
      </div>

      <div className="flex-1 flex items-center justify-center bg-slate-50 p-8 md:p-16">
        <Card className="w-full max-w-md">
          <div className="text-center mb-2">
            <h1 className="text-6xl font-bold mb-6 bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
              ¡Hola!
            </h1>
            <p className="text-base text-current-line font-normal">¿Ya tienes cuenta? <a href="/login" className="text-blue-800 hover:underline">Inicia sesión</a></p>
          </div>

          <h1 className="text-2xl text-current-line font-bold mb-4 mt-4">Crear Cuenta</h1>
          
          <form className="space-y-4" onSubmit={registerUser}>
            <div>
              <label htmlFor="name" className="block text-sm font-medium text-current-line mb-1">
                Nombre
              </label>
              <input
                type="text"
                id="name"
                name="name"
                className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
                placeholder="Tu nombre"
                value={credentials.name}
                onChange={changeUser}
              />
              {nameError && <p className="text-error text-sm mt-1">{nameError}</p>}
            </div>
            <div>
              <label htmlFor="email" className="block text-sm font-medium text-current-line mb-1">
                Email
              </label>
              <input
                type="email"
                id="email"
                name="email"
                className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
                placeholder="tu@email.com"
                value={credentials.email}
                onChange={changeUser}
              />
              {emailError && <p className="text-error text-sm mt-1">{emailError}</p>}
            </div>
            <div>
              <label htmlFor="password" className="block text-sm font-medium text-current-line mb-1">
                Contraseña
              </label>
              <div className="relative">
                <input
                  type={showPassword ? "text" : "password"}
                  id="password"
                  name="password"
                  className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
                  placeholder="••••••••"
                  value={credentials.password}
                  onChange={changeUser}
                />
                <button
                  type="button"
                  onClick={() => setShowPassword(!showPassword)}
                  className="absolute inset-y-0 right-0 pr-3 flex items-center text-gray-500"
                >
                  {showPassword ? <FaEyeSlash /> : <FaEye />}
                </button>
              </div>
              {passwordError && <p className="text-error text-sm mt-1">{passwordError}</p>}
            </div>
            <div>
              <label htmlFor="avatar" className="block text-sm font-medium text-current-line mb-1">
                Foto de perfil
              </label>
              <input
                type="file"
                id="avatar"
                name="avatar"
                className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
                onChange={handleFileChange}
              />
            </div>
            <Button className="w-full btn-primary" type="submit">
              Crear Cuenta
            </Button>
          </form>
        </Card>
      </div>
    </div>
  );
}
