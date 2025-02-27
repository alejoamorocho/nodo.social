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
    <Link
      href={`/nodes/${node.id}`}
      className="group block hover:scale-105 transform transition-transform duration-300"
    >
      <article className="bg-current-line/10 rounded-lg hover:bg-current-line/20 transition-colors border border-current-line/20 shadow-lg hover:shadow-xl overflow-hidden">
        {/* Image Section (Top of the card) */}
        <div className="flex justify-center">
          <div className="relative w-full aspect-video">
            <Image
              src={node.imageUrl}
              alt={node.name || 'Imagen del nodo'}
              fill
              className="object-cover rounded-t-lg"
            />
            {/* Gradient Overlay */}
            <div className="absolute inset-0 bg-gradient-to-t from-background/80 to-transparent" />
          </div>
        </div>

        {/* Category Badge */}
        <span
          className={`block mt-4 mx-auto px-3 py-1 text-sm rounded-full ${
            node.category === 'social'
              ? 'bg-purple/20 text-purple'
              : node.category === 'ambiental'
              ? 'bg-green/20 text-green'
              : 'bg-pink/20 text-pink'
          }`}
        >
          {node.category}
        </span>

        {/* Content Section (Bottom of the card) */}
        <div className="p-6 space-y-4">
          {/* Node Name */}
          <h3 className="text-2xl font-bold group-hover:text-primary transition-colors">
            {node.name}
          </h3>

          {/* Description */}
          <p className="text-foreground/80 line-clamp-2">
            {node.description}
          </p>

          {/* Location */}
          <div className="flex items-center gap-2 text-foreground/70">
            <MapPin className="w-5 h-5" />
            <span>{node.location || 'Ubicación no disponible'}</span>
          </div>

          {/* Products Preview */}
          {node.products && node.products.length > 0 && (
            <div className="mt-4">
              <div className="flex items-center gap-2 text-sm text-purple-400">
                <Tag className="w-4 h-4" />
                <span>Productos:</span>
              </div>
              <ul className="list-disc list-inside mt-2">
                {node.products.map((product, index) => (
                  <li key={index} className="text-foreground/80">
                    {product.name} - ${product.price}
                  </li>
                ))}
              </ul>
            </div>
          )}

          {/* Tags */}
          <div className="flex flex-wrap gap-2 mt-4">
            {node.tags?.slice(0, 3).map((tag, index) => (
              <span
                key={index}
                className="bg-background px-2 py-1 rounded-full text-xs text-foreground/70"
              >
                {tag}
              </span>
            ))}
            {node.tags && node.tags.length > 3 && (
              <span className="bg-background px-2 py-1 rounded-full text-xs text-foreground/70">
                +{node.tags.length - 3}
              </span>
            )}
          </div>

          {/* Created By and Date */}
          <div className="mt-4 text-sm text-foreground/70">
            <p>
              <span className="font-semibold">Creado por:</span> {node.createdBy}
            </p>
            <p>
              <span className="font-semibold">Fecha de creación:</span>{' '}
              {new Date(node.createdAt).toLocaleDateString()}
            </p>
          </div>
        </div>
      </article>
    </Link>
  )
}