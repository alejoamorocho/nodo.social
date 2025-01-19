'use client'

import { useState } from 'react'
import { Button } from '@/components/ui/Button'
import { Search, Filter, ShoppingCart } from 'lucide-react'

export default function StorePage() {
  const [searchQuery, setSearchQuery] = useState('')

  return (
    <div>
      <div className="flex flex-col md:flex-row justify-between items-center mb-8 gap-4">
        <h1 className="text-3xl font-bold">Tienda Social</h1>
        
        <div className="flex w-full md:w-auto gap-4">
          <div className="relative flex-1 md:w-80">
            <input
              type="text"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              placeholder="Buscar productos..."
              className="input pl-10 pr-4 py-2 w-full"
            />
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-comment" />
          </div>
          
          <Button variant="secondary">
            <Filter className="w-5 h-5 mr-2" />
            Filtros
          </Button>

          <Button variant="ghost" className="relative">
            <ShoppingCart className="w-5 h-5" />
            <span className="absolute -top-1 -right-1 w-5 h-5 rounded-full bg-primary text-xs flex items-center justify-center">
              0
            </span>
          </Button>
        </div>
      </div>

      {/* Grid de Productos */}
      <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6">
        {/* Placeholder para productos */}
        {Array.from({ length: 8 }).map((_, index) => (
          <div
            key={index}
            className="bg-current-line rounded-lg overflow-hidden hover:bg-current-line/80 transition-colors"
          >
            <div className="aspect-square bg-gradient-to-br from-purple/20 to-cyan/20" />
            
            <div className="p-4">
              <h3 className="text-lg font-semibold mb-2">Producto {index + 1}</h3>
              <p className="text-comment text-sm mb-4">
                Descripción breve del producto y sus características principales.
              </p>
              
              <div className="flex flex-wrap gap-2 mb-4">
                <span className="px-2 py-1 text-xs rounded bg-purple/20 text-purple">
                  Categoría
                </span>
                <span className="px-2 py-1 text-xs rounded bg-green/20 text-green">
                  Eco-friendly
                </span>
              </div>

              <div className="flex justify-between items-center">
                <span className="font-semibold">
                  $99.99
                </span>
                <Button variant="primary" size="sm">
                  Agregar
                </Button>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}
