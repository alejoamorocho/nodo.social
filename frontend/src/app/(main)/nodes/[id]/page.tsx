'use client'

import { useState } from 'react'
import { Button } from '@/components/ui/Button'
import { UpdateCard } from '@/components/fundraising/UpdateCard'
import { MapPin, Users, Heart, Share2, Flag } from 'lucide-react'

// Datos de ejemplo
const node = {
  id: 1,
  title: 'Refugio Animal Sustentable',
  description: `Estamos creando un espacio seguro y autosustentable para animales rescatados en la Ciudad de México. 
  Nuestro objetivo es proporcionar un hogar temporal a más de 200 animales mientras encontramos familias permanentes para ellos.
  
  El refugio contará con:
  - Área de recuperación médica
  - Espacios verdes para ejercicio
  - Zona de socialización
  - Huerto para alimentos frescos
  - Sistema de energía solar
  
  Tu apoyo nos ayudará a:
  1. Construir las instalaciones básicas
  2. Equipar el área médica
  3. Implementar sistemas sustentables
  4. Crear programas de adopción responsable`,
  category: 'animal',
  location: 'Ciudad de México',
  creator: {
    name: 'María González',
    avatar: '/avatar.jpg',
    bio: 'Veterinaria y activista por los derechos animales',
    projectsCreated: 3,
    projectsBacked: 12
  },
  updates: [
    {
      title: '¡Nuevo avance en el proyecto!',
      content: 'Hemos comenzado con la primera fase de construcción.',
      date: new Date('2024-01-05'),
      author: {
        name: 'María González',
        avatar: '/avatar.jpg'
      },
      likes: 45,
      comments: 12
    }
  ],
  products: [
    {
      id: 1,
      title: 'Camiseta Solidaria',
      description: 'Camiseta 100% algodón orgánico',
      imageUrl: '/product1.jpg',
      price: 299
    },
    {
      id: 2,
      title: 'Taza Ecológica',
      description: 'Taza reutilizable de bambú',
      imageUrl: '/product2.jpg',
      price: 199
    }
  ]
}

export default function NodePage({ params }: { params: { id: string } }) {
  const [activeTab, setActiveTab] = useState<'about' | 'updates' | 'products'>('about')

  return (
    <div className="max-w-6xl mx-auto">
      {/* Hero */}
      <div className="aspect-video bg-gradient-to-br from-purple/20 to-cyan/20 rounded-lg mb-8">
        {/* Aquí iría la imagen o video principal */}
      </div>

      {/* Content Grid */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        {/* Main Content */}
        <div className="lg:col-span-2 space-y-8">
          {/* Header */}
          <div>
            <div className="flex items-center gap-4 text-comment mb-4">
              <span className="px-3 py-1 rounded-full bg-purple/20 text-purple">
                Proyecto Animal
              </span>
              <span className="flex items-center gap-1">
                <MapPin className="w-4 h-4" />
                {node.location}
              </span>
            </div>

            <h1 className="text-4xl font-bold mb-4">{node.title}</h1>
            
            <div className="flex items-center gap-6">
              <Button variant="primary" size="lg">
                Contactar
              </Button>
              <Button variant="outline" size="lg">
                <Heart className="w-5 h-5" />
              </Button>
              <Button variant="outline" size="lg">
                <Share2 className="w-5 h-5" />
              </Button>
            </div>
          </div>

          {/* Tabs */}
          <div className="border-b border-current-line">
            <div className="flex gap-8">
              <button
                onClick={() => setActiveTab('about')}
                className={`pb-4 relative ${
                  activeTab === 'about'
                    ? 'text-foreground'
                    : 'text-comment hover:text-foreground'
                }`}
              >
                Acerca del Proyecto
                {activeTab === 'about' && (
                  <div className="absolute bottom-0 left-0 right-0 h-0.5 bg-purple" />
                )}
              </button>
              <button
                onClick={() => setActiveTab('updates')}
                className={`pb-4 relative ${
                  activeTab === 'updates'
                    ? 'text-foreground'
                    : 'text-comment hover:text-foreground'
                }`}
              >
                Actualizaciones
                {activeTab === 'updates' && (
                  <div className="absolute bottom-0 left-0 right-0 h-0.5 bg-purple" />
                )}
              </button>
              <button
                onClick={() => setActiveTab('products')}
                className={`pb-4 relative ${
                  activeTab === 'products'
                    ? 'text-foreground'
                    : 'text-comment hover:text-foreground'
                }`}
              >
                Productos
                {activeTab === 'products' && (
                  <div className="absolute bottom-0 left-0 right-0 h-0.5 bg-purple" />
                )}
              </button>
            </div>
          </div>

          {/* Tab Content */}
          {activeTab === 'about' && (
            <div className="prose prose-invert max-w-none">
              <p className="whitespace-pre-line">{node.description}</p>
            </div>
          )}

          {activeTab === 'updates' && (
            <div className="space-y-6">
              {node.updates.map((update, index) => (
                <UpdateCard key={index} update={update} />
              ))}
            </div>
          )}

          {activeTab === 'products' && (
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              {node.products.map((product) => (
                <div key={product.id} className="bg-current-line rounded-lg p-4">
                  <div className="aspect-square bg-background rounded-lg mb-4">
                    {/* Product image */}
                  </div>
                  <h3 className="text-lg font-bold mb-2">{product.title}</h3>
                  <p className="text-comment mb-4">{product.description}</p>
                  <div className="flex items-center justify-between">
                    <span className="text-xl font-bold">${product.price}</span>
                    <Button variant="primary" size="sm">Ver Detalles</Button>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>

        {/* Sidebar */}
        <div className="space-y-6">
          {/* Creator Card */}
          <div className="bg-current-line rounded-lg p-6">
            <h3 className="text-lg font-bold mb-4">Creador del Proyecto</h3>
            <div className="flex items-center gap-4 mb-4">
              <div className="w-12 h-12 bg-background rounded-full">
                {/* Creator avatar */}
              </div>
              <div>
                <h4 className="font-bold">{node.creator.name}</h4>
                <p className="text-sm text-comment">{node.creator.projectsCreated} proyectos creados</p>
              </div>
            </div>
            <p className="text-comment mb-4">{node.creator.bio}</p>
            <Button variant="outline" className="w-full">
              Ver Perfil
            </Button>
          </div>

          {/* Report Button */}
          <Button variant="ghost" className="w-full text-comment">
            <Flag className="w-4 h-4 mr-2" />
            Reportar Proyecto
          </Button>
        </div>
      </div>
    </div>
  )
}
