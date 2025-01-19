'use client'

import Image from 'next/image'
import Link from 'next/link'
import { MapPin, Tag } from 'lucide-react'
import type { Node } from '@/domain/models/Node'

interface NodeCardProps {
  node: Node
}

export function NodeCard({ node }: NodeCardProps) {
  return (
    <Link href={`/nodes/${node.id}`}>
      <div className="group bg-current-line rounded-lg overflow-hidden transition-transform hover:scale-[1.02]">
        {/* Image */}
        <div className="relative aspect-video">
          <Image
            src={node.imageUrl}
            alt={node.name}
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
              <span>{node.location}</span>
            </div>
          </div>

          <p className="text-comment line-clamp-2">{node.description}</p>

          {/* Products Preview */}
          {node.products.length > 0 && (
            <div className="space-y-2">
              <div className="flex items-center gap-2 text-sm text-purple-400">
                <Tag className="w-4 h-4" />
                <span>{node.products.length} productos vinculados</span>
              </div>
              <div className="flex -space-x-2">
                {node.products.slice(0, 3).map((product) => (
                  <div
                    key={product.id}
                    className="relative w-8 h-8 rounded-full overflow-hidden border-2 border-background"
                  >
                    <Image
                      src={product.imageUrl}
                      alt={product.title}
                      fill
                      className="object-cover"
                    />
                  </div>
                ))}
                {node.products.length > 3 && (
                  <div className="relative w-8 h-8 rounded-full bg-purple-500 border-2 border-background flex items-center justify-center text-xs">
                    +{node.products.length - 3}
                  </div>
                )}
              </div>
            </div>
          )}

          {/* Tags */}
          <div className="flex flex-wrap gap-2">
            {node.tags.slice(0, 3).map((tag, index) => (
              <span
                key={index}
                className="bg-background px-2 py-1 rounded-full text-xs"
              >
                {tag}
              </span>
            ))}
            {node.tags.length > 3 && (
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
