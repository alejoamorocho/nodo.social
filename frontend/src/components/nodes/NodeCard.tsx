'use client'

import React from 'react'
import Image from 'next/image'
import Link from 'next/link'
import { MapPin, Tag } from 'lucide-react'

interface Product {
  name: string
  price: number
  description: string
}

interface Node {
  id: string
  name: string
  description: string
  category: string
  tags: string[]
  imageUrl: string
  createdAt: Date
  createdBy: string
  products?: Product[]
  location?: string
}

interface NodeCardProps {
  node: Node
}

export const NodeCard: React.FC<NodeCardProps> = ({ node }) => {
  return (
    <Link href={`/nodes/${node.id}`}>
      <div className="group bg-current-line rounded-lg overflow-hidden transition-transform hover:scale-[1.02]">
        {/* Image */}
        <div className="relative aspect-video">
          <Image
            src={node.imageUrl}
            alt={node.name || 'Imagen del nodo'}
            fill
            className="object-cover"
          />
          <div className="absolute inset-0 bg-gradient-to-t from-background/80 to-transparent" />
        </div>

        {/* Content */}
        <div className="p-4 space-y-4">
          <div>
            <h3 className="text-xl font-bold group-hover:text-purple-400 transition-colors">
              {node.name}
            </h3>
            <div className="flex items-center gap-2 text-comment mt-1">
              <MapPin className="w-4 h-4" />
              <span>{node.location || 'Ubicación no disponible'}</span>
            </div>
          </div>

          <p className="text-comment line-clamp-2">{node.description}</p>
          <p className="text-comment"><strong>Categoría:</strong> {node.category}</p>
          <p className="text-comment"><strong>Creado por:</strong> {node.createdBy}</p>
          <p className="text-comment"><strong>Fecha de creación:</strong> {new Date(node.createdAt).toLocaleDateString()}</p>

          {/* Products Preview */}
          {node.products && node.products.length > 0 && (
            <div className="space-y-2">
              <div className="flex items-center gap-2 text-sm text-purple-400">
                <Tag className="w-4 h-4" />
                <span>Productos:</span>
              </div>
              <ul className="list-disc list-inside">
                {node.products.map((product, index) => (
                  <li key={index}>{product.name} - ${product.price}</li>
                ))}
              </ul>
            </div>
          )}

          {/* Tags */}
          <div className="flex flex-wrap gap-2">
            {node.tags?.slice(0, 3).map((tag, index) => (
              <span
                key={index}
                className="bg-background px-2 py-1 rounded-full text-xs"
              >
                {tag}
              </span>
            ))}
            {node.tags && node.tags.length > 3 && (
              <span className="bg-background px-2 py-1 rounded-full text-xs">
                +{node.tags.length - 3}
              </span>
            )}
          </div>
        </div>
      </div>
    </Link>
  )
}