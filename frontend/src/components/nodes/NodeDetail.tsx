'use client'

import Image from 'next/image'
import { MapPin, Link as LinkIcon } from 'lucide-react'
import { SocialMetrics } from '../social/SocialMetrics'
import { ProductCard } from '../products/ProductCard'
import type { Node } from '@/domain/models/Node'

interface NodeDetailProps {
  node: Node
}

export function NodeDetail({ node }: NodeDetailProps) {
  return (
    <div className="space-y-8">
      {/* Hero Section */}
      <div className="relative h-64 rounded-xl overflow-hidden">
        <Image
          src={node.imageUrl}
          alt={node.name}
          fill
          className="object-cover"
        />
        <div className="absolute inset-0 bg-gradient-to-t from-background/80 to-transparent" />
        <div className="absolute bottom-0 left-0 right-0 p-6">
          <h1 className="text-3xl font-bold">{node.name}</h1>
          <div className="flex items-center gap-2 text-comment mt-2">
            <MapPin className="w-4 h-4" />
            <span>{node.location}</span>
          </div>
        </div>
      </div>

      {/* Social Metrics */}
      <SocialMetrics
        connections={node.socialMetrics.connections}
        shares={node.socialMetrics.shares}
        points={node.socialMetrics.points}
        comments={node.socialMetrics.comments}
      />

      {/* Description */}
      <div className="prose prose-invert max-w-none">
        <p>{node.description}</p>
      </div>

      {/* Tags */}
      <div className="flex flex-wrap gap-2">
        {node.tags.map((tag, index) => (
          <span
            key={index}
            className="bg-current-line px-3 py-1 rounded-full text-sm"
          >
            {tag}
          </span>
        ))}
      </div>

      {/* Contact Information */}
      <div className="bg-current-line rounded-lg p-6 space-y-4">
        <h2 className="text-xl font-bold">Contacto</h2>
        <div className="space-y-2">
          {node.contacts.instagram && (
            <a
              href={`https://instagram.com/${node.contacts.instagram}`}
              target="_blank"
              rel="noopener noreferrer"
              className="flex items-center gap-2 text-pink-500 hover:text-pink-600 transition-colors"
            >
              <LinkIcon className="w-4 h-4" />
              <span>@{node.contacts.instagram}</span>
            </a>
          )}
          {node.contacts.whatsapp && (
            <a
              href={`https://wa.me/${node.contacts.whatsapp}`}
              target="_blank"
              rel="noopener noreferrer"
              className="flex items-center gap-2 text-green-500 hover:text-green-600 transition-colors"
            >
              <LinkIcon className="w-4 h-4" />
              <span>WhatsApp</span>
            </a>
          )}
          <a
            href={`mailto:${node.contacts.email}`}
            className="flex items-center gap-2 text-blue-500 hover:text-blue-600 transition-colors"
          >
            <LinkIcon className="w-4 h-4" />
            <span>{node.contacts.email}</span>
          </a>
        </div>
      </div>

      {/* Associated Products */}
      <div className="space-y-6">
        <h2 className="text-2xl font-bold">Productos Vinculados</h2>
        <p className="text-comment">
          Al comprar estos productos, un porcentaje va directamente a apoyar este nodo
        </p>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {node.products.map((product) => (
            <ProductCard
              key={product.id}
              title={product.title}
              description={product.description}
              imageUrl={product.imageUrl}
              nodeContribution={product.nodeContribution}
              instagramHandle={product.contacts.instagram}
              whatsappNumber={product.contacts.whatsapp}
              tags={product.tags}
            />
          ))}
        </div>
      </div>
    </div>
  )
}
