'use client'

import Image from 'next/image'
import { Instagram, WhatsApp } from 'lucide-react'

interface ProductCardProps {
  title: string
  description: string
  imageUrl: string
  nodeContribution: number
  instagramHandle?: string
  whatsappNumber?: string
  tags: string[]
}

export function ProductCard({
  title,
  description,
  imageUrl,
  nodeContribution,
  instagramHandle,
  whatsappNumber,
  tags
}: ProductCardProps) {
  return (
    <div className="rounded-lg bg-current-line p-4 space-y-4">
      {/* Product Image */}
      <div className="relative aspect-square rounded-lg overflow-hidden">
        <Image
          src={imageUrl}
          alt={title}
          fill
          className="object-cover"
        />
      </div>

      {/* Product Info */}
      <div className="space-y-2">
        <h3 className="text-xl font-bold">{title}</h3>
        <p className="text-comment">{description}</p>
        
        {/* Node Contribution */}
        <div className="bg-purple-500/20 text-purple-500 px-3 py-1 rounded-full text-sm inline-block">
          {nodeContribution}% para el nodo
        </div>

        {/* Tags */}
        <div className="flex flex-wrap gap-2">
          {tags.map((tag, index) => (
            <span
              key={index}
              className="bg-background px-2 py-1 rounded-full text-xs"
            >
              {tag}
            </span>
          ))}
        </div>
      </div>

      {/* Contact Links */}
      <div className="flex gap-4">
        {instagramHandle && (
          <a
            href={`https://instagram.com/${instagramHandle}`}
            target="_blank"
            rel="noopener noreferrer"
            className="flex items-center gap-2 text-pink-500 hover:text-pink-600 transition-colors"
          >
            <Instagram className="w-5 h-5" />
            <span>@{instagramHandle}</span>
          </a>
        )}
        
        {whatsappNumber && (
          <a
            href={`https://wa.me/${whatsappNumber}`}
            target="_blank"
            rel="noopener noreferrer"
            className="flex items-center gap-2 text-green-500 hover:text-green-600 transition-colors"
          >
            <WhatsApp className="w-5 h-5" />
            <span>WhatsApp</span>
          </a>
        )}
      </div>
    </div>
  )
}
