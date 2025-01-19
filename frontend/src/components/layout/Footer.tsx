'use client'

import Link from 'next/link'
import { Twitter, Instagram, Linkedin, Github, Heart } from 'lucide-react'

export function Footer() {
  return (
    <footer className="bg-background border-t border-current-line">
      <div className="container mx-auto px-4 py-12">
        {/* Main Footer Content */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-8 mb-8">
          {/* Brand Section */}
          <div className="space-y-4">
            <Link href="/" className="block">
              <span className="text-2xl font-bold bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
                NODO.SOCIAL
              </span>
            </Link>
            <p className="text-comment">
              Conectando personas y emprendedores con causas sociales, ambientales y animales.
            </p>
          </div>

          {/* Quick Links */}
          <div>
            <h3 className="text-lg font-semibold text-foreground mb-4">Explorar</h3>
            <ul className="space-y-2">
              <li>
                <Link href="/nodes" className="text-comment hover:text-primary transition-colors">
                  Nodos
                </Link>
              </li>
              <li>
                <Link href="/store" className="text-comment hover:text-primary transition-colors">
                  Tienda
                </Link>
              </li>
              <li>
                <Link href="/feed" className="text-comment hover:text-primary transition-colors">
                  Feed
                </Link>
              </li>
            </ul>
          </div>

          {/* Resources */}
          <div>
            <h3 className="text-lg font-semibold text-foreground mb-4">Recursos</h3>
            <ul className="space-y-2">
              <li>
                <Link href="/about" className="text-comment hover:text-primary transition-colors">
                  Acerca de
                </Link>
              </li>
              <li>
                <Link href="/contact" className="text-comment hover:text-primary transition-colors">
                  Contacto
                </Link>
              </li>
              <li>
                <Link href="/faq" className="text-comment hover:text-primary transition-colors">
                  FAQ
                </Link>
              </li>
            </ul>
          </div>

          {/* Legal */}
          <div>
            <h3 className="text-lg font-semibold text-foreground mb-4">Legal</h3>
            <ul className="space-y-2">
              <li>
                <Link href="/privacy" className="text-comment hover:text-primary transition-colors">
                  Privacidad
                </Link>
              </li>
              <li>
                <Link href="/terms" className="text-comment hover:text-primary transition-colors">
                  Términos
                </Link>
              </li>
              <li>
                <Link href="/cookies" className="text-comment hover:text-primary transition-colors">
                  Cookies
                </Link>
              </li>
            </ul>
          </div>
        </div>

        {/* Social Links */}
        <div className="flex flex-wrap justify-center gap-6 py-8 border-t border-current-line">
          <a
            href="https://twitter.com"
            target="_blank"
            rel="noopener noreferrer"
            className="text-comment hover:text-primary transition-colors"
            aria-label="Síguenos en Twitter"
          >
            <Twitter className="w-6 h-6" aria-hidden="true" />
            <span className="sr-only">Twitter</span>
          </a>
          <a
            href="https://instagram.com"
            target="_blank"
            rel="noopener noreferrer"
            className="text-comment hover:text-pink transition-colors"
            aria-label="Síguenos en Instagram"
          >
            <Instagram className="w-6 h-6" aria-hidden="true" />
            <span className="sr-only">Instagram</span>
          </a>
          <a
            href="https://linkedin.com"
            target="_blank"
            rel="noopener noreferrer"
            className="text-comment hover:text-cyan transition-colors"
            aria-label="Síguenos en LinkedIn"
          >
            <Linkedin className="w-6 h-6" aria-hidden="true" />
            <span className="sr-only">LinkedIn</span>
          </a>
          <a
            href="https://github.com"
            target="_blank"
            rel="noopener noreferrer"
            className="text-comment hover:text-purple transition-colors"
            aria-label="Visita nuestro GitHub"
          >
            <Github className="w-6 h-6" aria-hidden="true" />
            <span className="sr-only">GitHub</span>
          </a>
        </div>

        {/* Copyright */}
        <div className="text-center text-comment pt-8 border-t border-current-line">
          <p className="flex items-center justify-center gap-1">
            &copy; {new Date().getFullYear()} NODO.SOCIAL. Hecho con 
            <Heart className="w-4 h-4 text-red fill-red" aria-hidden="true" /> 
            en México
          </p>
        </div>
      </div>
    </footer>
  )
}
