'use client'

import Link from 'next/link'
import { Button } from '@/components/ui/Button'
import { Menu, Bell, User, Search, LogOut } from 'lucide-react'
import { useState, useCallback, useEffect } from 'react'
import { onAuthStateChanged, signOut } from 'firebase/auth'
import { auth } from '../../firebase'
import Image from 'next/image'

interface User {
  name: string
  email: string
  avatar: string
}

export function Navbar() {
  const [isMenuOpen, setIsMenuOpen] = useState(false)
  const [user, setUser] = useState<User | null>(null)

  const toggleMenu = useCallback(() => {
    setIsMenuOpen((prev) => !prev)
  }, [])

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, (currentUser) => {
      if (currentUser) {
        setUser({
          name: currentUser.displayName || '',
          email: currentUser.email || '',
          avatar: currentUser.photoURL || '',
        })
      } else {
        setUser(null)
      }
    })

    return () => unsubscribe()
  }, [])

  const handleSignOut = async () => {
    await signOut(auth)
    setUser(null)
  }

  return (
    <nav className="sticky top-0 z-50 bg-background/80 backdrop-blur-sm border-b border-current-line">
      <div className="container mx-auto px-4">
        <div className="flex items-center justify-between h-16">
          {/* Logo */}
          <Link href="/" className="flex items-center space-x-2">
            <span className="text-2xl font-bold bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
              NODO.SOCIAL
            </span>
          </Link>

          {/* Desktop Navigation */}
          <div className="hidden md:flex items-center space-x-1">
            <div className="relative">
              <input
                type="text"
                placeholder="Buscar..."
                className="input pl-10 pr-4 py-1.5 text-sm w-64"
                aria-label="Buscar en la plataforma"
              />
              <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-comment" aria-hidden="true" />
            </div>
            <Link href="/nodes">
              <Button variant="ghost" size="sm">Nodos</Button>
            </Link>
            <Link href="/store">
              <Button variant="ghost" size="sm">Tienda</Button>
            </Link>
            <Link href="/feed">
              <Button variant="ghost" size="sm">Feed</Button>
            </Link>
          </div>

          {/* User Actions */}
          <div className="hidden md:flex items-center space-x-2">
            <Button variant="ghost" size="sm" aria-label="Notificaciones">
              <Bell className="w-5 h-5" aria-hidden="true" />
            </Button>
            {user ? (
              <div className="relative">
                <Button variant="ghost" size="sm" onClick={toggleMenu} aria-expanded={isMenuOpen} aria-controls="user-menu" aria-label="Menú de usuario" className="border border-primary">
                  {user.avatar ? (
                    <Image
                      src={user.avatar}
                      alt={user.name}
                      width={32}
                      height={32}
                      className="rounded-full"
                    />
                  ) : (
                    <User className="w-5 h-5" aria-hidden="true" />
                  )}
                  <span className="ml-2">{user.name}</span>
                </Button>
                {isMenuOpen && (
                  <div id="user-menu" className="absolute right-0 mt-2 w-48 bg-white border border-gray-200 rounded-md shadow-lg py-1 z-20">
                    <Link href="/profile" className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
                      Ver Perfil
                    </Link>
                    <button
                      onClick={handleSignOut}
                      className="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                    >
                      Cerrar Sesión
                    </button>
                  </div>
                )}
              </div>
            ) : (
              <Link href="/login">
                <Button variant="primary" size="sm">
                  <User className="w-5 h-5 mr-2" aria-hidden="true" />
                  Iniciar Sesión
                </Button>
              </Link>
            )}
          </div>

          {/* Mobile Menu Button */}
          <Button
            variant="ghost"
            size="sm"
            className="md:hidden"
            onClick={toggleMenu}
            aria-expanded={isMenuOpen}
            aria-controls="mobile-menu"
            aria-label="Menú de navegación"
          >
            <Menu className="w-6 h-6" aria-hidden="true" />
          </Button>
        </div>

        {/* Mobile Menu */}
        <div 
          id="mobile-menu"
          className={`md:hidden py-4 border-t border-current-line ${isMenuOpen ? 'block' : 'hidden'}`}
          role="navigation"
          aria-label="Menú móvil"
        >
          <div className="space-y-4">
            <div className="relative">
              <input
                type="text"
                placeholder="Buscar..."
                className="input pl-10 pr-4 py-2 w-full"
                aria-label="Buscar en la plataforma"
              />
              <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-comment" aria-hidden="true" />
            </div>
            <div className="grid grid-cols-1 gap-2">
              <Link href="/nodes">
                <Button variant="ghost" className="w-full justify-start">
                  Nodos
                </Button>
              </Link>
              <Link href="/store">
                <Button variant="ghost" className="w-full justify-start">
                  Tienda
                </Button>
              </Link>
              <Link href="/feed">
                <Button variant="ghost" className="w-full justify-start">
                  Feed
                </Button>
              </Link>
              {user ? (
                <>
                  <Link href="/profile">
                    <Button variant="primary" className="w-full justify-start">
                      <User className="w-5 h-5 mr-2" aria-hidden="true" />
                      Ver Perfil
                    </Button>
                  </Link>
                  <Button
                    variant="ghost"
                    className="w-full justify-start"
                    onClick={handleSignOut}
                  >
                    <LogOut className="w-5 h-5 mr-2" aria-hidden="true" />
                    Cerrar Sesión
                  </Button>
                </>
              ) : (
                <Link href="/login">
                  <Button variant="primary" className="w-full justify-start">
                    <User className="w-5 h-5 mr-2" aria-hidden="true" />
                    Iniciar Sesión
                  </Button>
                </Link>
              )}
            </div>
          </div>
        </div>
      </div>
    </nav>
  )
}
