'use client'

import { Users, Share2, Star, MessageSquare } from 'lucide-react'

interface SocialMetricsProps {
  connections: number
  shares: number
  points: number
  comments: number
}

export function SocialMetrics({ connections, shares, points, comments }: SocialMetricsProps) {
  return (
    <div className="space-y-4">
      {/* Metrics Grid */}
      <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
        <div className="flex items-center gap-2">
          <Users className="w-5 h-5 text-purple-500" />
          <div>
            <p className="text-2xl font-bold">{connections.toLocaleString()}</p>
            <p className="text-sm text-comment">Conexiones</p>
          </div>
        </div>
        
        <div className="flex items-center gap-2">
          <Share2 className="w-5 h-5 text-green-500" />
          <div>
            <p className="text-2xl font-bold">{shares.toLocaleString()}</p>
            <p className="text-sm text-comment">Compartidos</p>
          </div>
        </div>

        <div className="flex items-center gap-2">
          <Star className="w-5 h-5 text-yellow-500" />
          <div>
            <p className="text-2xl font-bold">{points.toLocaleString()}</p>
            <p className="text-sm text-comment">Puntos</p>
          </div>
        </div>

        <div className="flex items-center gap-2">
          <MessageSquare className="w-5 h-5 text-blue-500" />
          <div>
            <p className="text-2xl font-bold">{comments.toLocaleString()}</p>
            <p className="text-sm text-comment">Comentarios</p>
          </div>
        </div>
      </div>
    </div>
  )
}
