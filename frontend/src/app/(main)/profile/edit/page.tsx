'use client'

import { useState, useEffect } from 'react'
import { useRouter } from 'next/navigation'
import { auth, db } from '../../../../firebase'
import { doc, getDoc, updateDoc } from 'firebase/firestore'
import { onAuthStateChanged, updatePassword, reauthenticateWithCredential, EmailAuthProvider } from 'firebase/auth'
import { getStorage, ref, uploadBytes, getDownloadURL } from 'firebase/storage'
import { Button } from '@/components/ui/Button'
import { FaEye, FaEyeSlash} from "react-icons/fa";

interface User {
  name: string
  email: string
  avatar: string
  coverImage: string
  bio: string
  location: string
  links: {
    website: string
    twitter: string
    instagram: string
  }
}

type NestedPartial<T> = {
  [P in keyof T]?: T[P] extends object ? NestedPartial<T[P]> : T[P];
};

export default function EditProfilePage() {
  const [user, setUser] = useState<User | null>(null)
  const [loading, setLoading] = useState(true)
  const [avatar, setAvatar] = useState<File | null>(null)
  const [newPassword, setNewPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [passwordError, setPasswordError] = useState('')
  const [currentPassword, setCurrentPassword] = useState('')
  const [showErrorPopup, setShowErrorPopup] = useState(false)
  const [showSuccessPopup, setShowSuccessPopup] = useState(false)
  const [showPassword, setShowPassword] = useState(false)
  const [showConfirmPassword, setShowConfirmPassword] = useState(false)
  const [showCurrentPassword, setShowCurrentPassword] = useState(false)
  const router = useRouter()

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, async (currentUser) => {
      if (currentUser) {
        try {
          const userDoc = await getDoc(doc(db, 'users', currentUser.uid))
          if (userDoc.exists()) {
            const userData = userDoc.data() as User
            console.log('Datos del usuario cargados:', userData)
            setUser(userData)
          } else {
            console.log('No se encontraron datos del usuario')
          }
        } catch (error) {
          console.error('Error al obtener los datos del usuario:', error)
        } finally {
          setLoading(false)
        }
      } else {
        router.push('/login')
      }
    })

    return () => unsubscribe()
  }, [router])

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target
    setUser(prevUser => prevUser ? { ...prevUser, [name]: value } : null)
  }

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      setAvatar(e.target.files[0])
      console.log('Archivo de avatar seleccionado:', e.target.files[0])
    }
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (user) {
      console.log('Datos del usuario antes de guardar:', user)
      try {
        const userRef = doc(db, 'users', auth.currentUser!.uid)
        if (avatar) {
          const storage = getStorage()
          const storageRef = ref(storage, `avatars/${auth.currentUser!.uid}`)
          await uploadBytes(storageRef, avatar)
          const avatarURL = await getDownloadURL(storageRef)
          console.log('URL del avatar subido:', avatarURL)
          user.avatar = avatarURL
        }
        await updateDoc(userRef, user as NestedPartial<User>)
        console.log('Datos del usuario actualizados en Firestore:', user)
        router.push('/profile')
      } catch (error) {
        console.error('Error al actualizar los datos del usuario:', error)
      }
    }
  }

  const handlePasswordChange = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (newPassword !== confirmPassword) {
      setPasswordError('Las contraseñas no coinciden');
      setShowErrorPopup(true);
      return;
    }
  
    try {
      const currentUser = auth.currentUser;
      if (currentUser && currentUser.email) {
        const credential = EmailAuthProvider.credential(currentUser.email, currentPassword);
        await reauthenticateWithCredential(currentUser, credential);
        await updatePassword(currentUser, newPassword);
        console.log('Contraseña actualizada correctamente');
        setNewPassword('');
        setConfirmPassword('');
        setCurrentPassword('');
        setPasswordError('');
        setShowSuccessPopup(true);
        setTimeout(() => {
          router.push('/profile');
        }, 2000); 
      } else {
        setPasswordError('No se pudo obtener la información del usuario.');
        setShowErrorPopup(true);
      }
    } catch (error) {
      if (error.code === 'auth/wrong-password') {
        setPasswordError('La contraseña actual es incorrecta');
      } else if (error.code === 'auth/invalid-credential') {
        setPasswordError('La contraseña actual no es válida');
      } else {
        setPasswordError('Error al actualizar la contraseña');
      }
      setShowErrorPopup(true);
    }
  };

  if (loading) {
    return <div>Loading...</div>
  }

  return (
    <div className="max-w-4xl mx-auto p-6">
      <h1 className="text-4xl font-bold mb-4 text-primary">Editar Perfil</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block text-sm font-medium">Nombre</label>
          <input
            type="text"
            name="name"
            value={user?.name || ''}
            onChange={handleChange}
            className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Email</label>
          <input
            type="email"
            name="email"
            value={user?.email || ''}
            onChange={handleChange}
            className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
            disabled
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Foto de perfil</label>
          <input
            type="file"
            name="avatar"
            className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
            onChange={handleFileChange}
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Biografía</label>
          <textarea
            name="bio"
            value={user?.bio || ''}
            onChange={handleChange}
            className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Ubicación</label>
          <input
            type="text"
            name="location"
            value={user?.location || ''}
            onChange={handleChange}
            className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Sitio web</label>
          <input
            type="text"
            name="links.website"
            value={user?.links.website || ''}
            onChange={handleChange}
            className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Twitter</label>
          <input
            type="text"
            name="links.twitter"
            value={user?.links.twitter || ''}
            onChange={handleChange}
            className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
          />
        </div>
        <div>
          <label className="block text-sm font-medium">Instagram</label>
          <input
            type="text"
            name="links.instagram"
            value={user?.links.instagram || ''}
            onChange={handleChange}
            className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
          />
        </div>
        <Button type="submit" variant="primary">Guardar Cambios</Button>
      </form>

      <h2 className="text-2xl font-bold mt-8 mb-4 text-primary">Cambiar Contraseña</h2>
      <form onSubmit={handlePasswordChange} className="space-y-4">
        <div className="relative">
          <label className="block text-sm font-medium">Contraseña Actual</label>
          <input
            type={showCurrentPassword ? 'text' : 'password'}
            name="currentPassword"
            value={currentPassword}
            onChange={(e) => setCurrentPassword(e.target.value)}
            className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
          />
          <button
            type="button"
            className="absolute inset-y-0 right-0 pr-3 flex items-center text-gray-500 mt-5"
            onClick={() => setShowCurrentPassword(!showCurrentPassword)}
          >
            {showCurrentPassword ? <FaEyeSlash /> : <FaEye />}
          </button>
        </div>
        <div className="relative">
          <label className="block text-sm font-medium">Nueva Contraseña</label>
          <input
            type={showPassword ? 'text' : 'password'}
            name="newPassword"
            value={newPassword}
            onChange={(e) => setNewPassword(e.target.value)}
            className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
          />
          <button
            type="button"
            className="absolute inset-y-0 right-0 pr-3 flex items-center text-gray-500 mt-5"
            onClick={() => setShowPassword(!showPassword)}
          >
            {showPassword ? <FaEyeSlash /> : <FaEye />}
          </button>
        </div>
        <div className="relative">
          <label className="block text-sm font-medium">Confirmar Nueva Contraseña</label>
          <input
            type={showConfirmPassword ? 'text' : 'password'}
            name="confirmPassword"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
          />
          <button
            type="button"
            className="absriitgray-500 mt0 riitgray-500 mtlex itgray-500 mt text-gray-500 mt-5"
            onClick={() => setShowConfirmPassword(!showConfirmPassword)}
          >
            {showPassword ? <FaEyeSlash /> : <FaEye />}
          </button>
        </div>
        {passwordError && <p className="text-red-500 text-sm">{passwordError}</p>}
        <Button type="submit" variant="primary">Actualizar Contraseña</Button>
      </form>

      {showErrorPopup && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
          <div className="bg-white p-6 rounded-lg shadow-lg">
            <h3 className="text-xl font-bold mb-4 text-black">Error</h3>
            <p className="text-black mb-4">{passwordError}</p>
            <Button onClick={() => setShowErrorPopup(false)} variant="secondary">Cerrar</Button>
          </div>
        </div>
      )}

      {showSuccessPopup && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
          <div className="bg-white p-6 rounded-lg shadow-lg">
            <h3 className="text-xl font-bold mb-4 text-black">Éxito</h3>
            <p className="text-black mb-4">Contraseña actualizada correctamente</p>
            <Button onClick={() => {
              setShowSuccessPopup(false);
              router.push('/profile');
            }} variant="secondary">Cerrar</Button>
          </div>
        </div>
      )}
    </div>
  )
}