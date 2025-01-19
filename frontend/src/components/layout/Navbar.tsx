'use client'

import Link from 'next/link'
import { Button } from '@/components/ui/Button'
import { Menu, Bell, User, Search } from 'lucide-react'
import { useState, useCallback } from 'react'

export function Navbar() {
  const [isMenuOpen, setIsMenuOpen] = useState(false)

  const toggleMenu = useCallback(() => {
    setIsMenuOpen((prev) => !prev)
  }, [])

  return (
    <nav className="sticky top-0 z-50 bg-background/80 backdrop-blur-sm border-b border-current-line">
      <div className="container mx-auto px-4">
        <div className="flex items-center justify-between h-16">
          {/* Logo */}
          <Link 
            href="/" 
            className="flex items-center space-x-2"
          >
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
            <Link href="/profile">
              <Button variant="primary" size="sm">
                <User className="w-5 h-5 mr-2" aria-hidden="true" />
                Perfil
              </Button>
            </Link>
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
              <Link href="/profile">
                <Button variant="primary" className="w-full justify-start">
                  <User className="w-5 h-5 mr-2" aria-hidden="true" />
                  Perfil
                </Button>
              </Link>
            </div>
          </div>
        </div>
      </div>
    </nav>
  )
}
