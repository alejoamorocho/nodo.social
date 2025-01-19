'use client'

import { Button } from '@/components/ui/Button'
import Link from 'next/link'

export default function HomePage() {
  return (
    <div className="space-y-8">
      <section className="text-center py-16 bg-gradient-to-r from-primary/10 to-secondary/10 rounded-lg">
        <h1 className="text-4xl font-bold mb-4 bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
          Bienvenido a NODO.SOCIAL
        </h1>
        <p className="text-xl text-foreground/80 max-w-2xl mx-auto">
          Conectando personas y emprendedores con causas sociales, ambientales y animales
          a través de nodos temáticos.
        </p>
      </section>

      <section className="grid md:grid-cols-3 gap-6">
        <div className="p-6 bg-background border border-current-line rounded-lg">
          <h2 className="text-xl font-semibold mb-3">Causas Sociales</h2>
          <p className="text-foreground/70">
            Descubre y apoya iniciativas que impactan positivamente en la sociedad.
          </p>
        </div>
        <div className="p-6 bg-background border border-current-line rounded-lg">
          <h2 className="text-xl font-semibold mb-3">Impacto Ambiental</h2>
          <p className="text-foreground/70">
            Únete a proyectos que protegen y preservan nuestro medio ambiente.
          </p>
        </div>
        <div className="p-6 bg-background border border-current-line rounded-lg">
          <h2 className="text-xl font-semibold mb-3">Bienestar Animal</h2>
          <p className="text-foreground/70">
            Colabora con iniciativas dedicadas al cuidado y protección animal.
          </p>
        </div>
      </section>

      <section className="flex justify-center gap-4">
        <Link href="/nodes/create">
          <Button variant="primary" size="lg">
            Crear un Nodo
          </Button>
        </Link>
        <Link href="/nodes">
          <Button variant="secondary" size="lg">
            Explorar Nodos
          </Button>
        </Link>
      </section>
    </div>
  )
}
