"use client";

import React, { useState } from "react";
import { Card } from "@/components/ui/Card";
import { Button } from "@/components/ui/Button";
import {
  signInWithEmailAndPassword,
  signInWithPopup,
  sendPasswordResetEmail,
} from "firebase/auth";
import { auth, googleProvider } from "../../../../firebase";
import { useRouter } from "next/navigation";
import { FaEye, FaEyeSlash, FaTimes } from "react-icons/fa";
import Image from "next/image";
import googleLogo from "../../../../public/google.png"; 

export default function LoginPage() {
  const [credentials, setCredentials] = useState({ email: "", password: "" });
  const [emailError, setEmailError] = useState("");
  const [passwordError, setPasswordError] = useState("");
  const [modalError, setModalError] = useState("");
  const [modalSuccess, setModalSuccess] = useState(""); 
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [loading, setLoading] = useState(false);
  const [showPassword, setShowPassword] = useState(false);
  const { push } = useRouter();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setCredentials({ ...credentials, [e.target.name]: e.target.value });
  };

  const validateEmail = (email: string) => {
    const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return re.test(String(email).toLowerCase());
  };

  const validatePassword = (password: string) => {
    return password.length >= 6;
  };

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    const { email, password } = credentials;

    // Reset errors
    setEmailError("");
    setPasswordError("");

    // Validar campos vacíos
    if (!email) {
      setEmailError("Por favor, ingresa tu email.");
    }
    if (!password) {
      setPasswordError("Por favor, ingresa tu contraseña.");
    }

    if (!email || !password) {
      return;
    }

    // Validar email
    if (!validateEmail(email)) {
      setEmailError("Por favor, ingresa un email válido.");
      return;
    }

    // Validar contraseña
    if (!validatePassword(password)) {
      setPasswordError("La contraseña debe tener al menos 6 caracteres.");
      return;
    }

    try {
      await signInWithEmailAndPassword(auth, email, password);
      push("/");
    } catch (error: unknown) {
      if (error instanceof Error) {
        setEmailError(error.message || "Hubo un problema al iniciar sesión.");
        console.error(error);
      }
    }
  };

  const handleGoogleLogin = async () => {
    setEmailError("");
    setPasswordError("");
    try {
      await signInWithPopup(auth, googleProvider);
      push("/");
    } catch (error: unknown) {
      if (error instanceof Error) {
        setEmailError(error.message || "Error al iniciar sesión con Google.");
        console.error(error);
      }
    }
  };

  const handlePasswordReset = async () => {
    if (!credentials.email) {
      setModalError("Ingresa tu email para recuperar la contraseña.");
      return;
    }
    setLoading(true);
    setModalError("");
    setModalSuccess("");

    try {
      await sendPasswordResetEmail(auth, credentials.email);
      setModalSuccess("Se envió un correo de recuperación.");
      setIsModalOpen(false);
    } catch (error: unknown) {
      if (error instanceof Error) {
        setModalError(error.message || "Error al enviar el correo de recuperación.");
        console.error(error);
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex flex-col md:flex-row">
      <div className="hidden md:flex flex-1 items-center justify-center p-8 md:p-16">
        <div className="text-center">
          <h1 className="text-3xl md:text-5xl font-bold mb-4 bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
            Bienvenido a NODO.SOCIAL
          </h1>
          <h2 className="text-xl md:text-xl lg:text-2xl font-semibold mb-32 mt-24">Conectando personas y emprendedores con causas sociales, ambientales y animales a través de nodos temáticos.</h2>
          <p className="text-lg md:text-base lg:text-lg mb-8">Accede a tu cuenta para seguir con tus actividades.</p>
        </div>
      </div>

      <div className="flex-1 flex items-center justify-center bg-slate-50 p-8 md:p-16">
        <Card className="w-full max-w-md">
          <div className="text-center mb-2">
            <h1 className="text-4xl md:text-6xl font-bold mb-4 bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
              Hola!
            </h1>
            <p className="text-base text-current-line font-normal m-11">¿No tienes cuenta? <a href="/register" className="text-blue-800 hover:underline">Regístrate</a></p>
          </div>

          <h1 className="text-2xl text-current-line font-bold mb-4 mt-4">Iniciar Sesión</h1>
          
          {emailError && <p className="text-error text-center mb-4">{emailError}</p>}
          {passwordError && <p className="text-error text-center mb-4">{passwordError}</p>}

          <form className="space-y-4" onSubmit={handleLogin}>
            <div>
              <label htmlFor="email" className="block text-base text-current-line font-medium">Email</label>
              <input
                type="email"
                id="email"
                name="email"
                placeholder="Ingresa tu@email.com"
                value={credentials.email}
                onChange={handleChange}
                required
                className="bg-gray-50 border-2 border-primary text-gray-900 text-sm rounded-lg focus:ring-blue-800 focus:border-blue-800 block w-full p-2.5"
              />
              {emailError && <p className="text-red-500 text-sm mt-1">{emailError}</p>}
            </div>
            <div>
              <label htmlFor="password" className="block text-base text-current-line">Contraseña</label>
              <div className="relative">
                <input
                  type={showPassword ? "text" : "password"}
                  id="password"
                  name="password"
                  placeholder="••••••••"
                  value={credentials.password}
                  onChange={handleChange}
                  required
                  className="w-full px-3 py-2 bg-gray-50 border-2 border-primary rounded-md focus:ring-2 text-gray-900"
                />
                <button
                  type="button"
                  onClick={() => setShowPassword(!showPassword)}
                  className="absolute inset-y-0 right-0 pr-3 flex items-center text-gray-500"
                >
                  {showPassword ? <FaEyeSlash /> : <FaEye />}
                </button>
              </div>
              {passwordError && <p className="text-red-500 text-sm mt-1">{passwordError}</p>}
            </div>
            <Button className="w-full btn-primary" type="submit">
              Iniciar Sesión
            </Button>
          </form>

          <button
            onClick={() => setIsModalOpen(true)}
            className="text-blue-800 text-sm mt-2"
          >
            ¿Olvidaste tu contraseña?
          </button>

          {/* Modal de recuperación de contraseña */}
          {isModalOpen && (
            <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
              <div className="bg-white p-8 rounded-lg shadow-lg w-3/4 md:w-1/2 lg:w-1/3">
                <div className="flex justify-between items-center mb-4">
                  <h2 className="text-xl font-bold text-primary">Recuperar Contraseña</h2>
                  <button onClick={() => setIsModalOpen(false)} className="text-primary">
                    <FaTimes />
                  </button>
                </div>
                <input
                  type="email"
                  name="email"
                  placeholder="Ingresa tu email"
                  value={credentials.email}
                  onChange={handleChange}
                  className="w-full px-3 py-2 border-2 border-stone-300 rounded-md mb-4 text-gray-900"
                />
                {modalError && <p className="text-error text-center text-sm mb-4">{modalError}</p>}
                <Button onClick={handlePasswordReset} disabled={loading} className="w-full btn-primary">
                  {loading ? "Enviando..." : "Recuperar contraseña"}
                </Button>
              </div>
            </div>
          )}

          {modalSuccess && (
            <div className="fixed top-0 inset-x-0 flex items-center justify-center transition-transform transform translate-y-0">
              <div className="bg-success text-current-line p-4 rounded-lg shadow-lg flex items-center">
                <p>{modalSuccess}</p>
                <button
                  onClick={() => setModalSuccess("")}
                  className="ml-4 text-current-line underline"
                >
                  <FaTimes />
                </button>
              </div>
            </div>
          )}

          <div className="mt-6 text-center">
            <button
              onClick={handleGoogleLogin}
              className="flex items-center justify-center gap-2 w-full py-2 border border-primary rounded-md transition text-primary hover:scale-110"
            >
              <Image src={googleLogo} alt="Google" className="h-6 w-6" />
              <span className="text-current-line">Regístrate con Google</span>
            </button>
          </div>
        </Card>
      </div>
    </div>
  );
}