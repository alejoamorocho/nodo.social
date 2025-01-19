'use client'

import { format } from 'date-fns'
import { es } from 'date-fns/locale'
import { MessageSquare, Heart } from 'lucide-react'

interface UpdateCardProps {
  title: string
  content: string
  date: Date
  author: {
    name: string
    avatar?: string
  }
  likes?: number
  comments?: number
  onLike?: () => void
  onComment?: () => void
}

export function UpdateCard({
  title,
  content,
  date,
  author,
  likes = 0,
  comments = 0,
  onLike,
  onComment
}: UpdateCardProps) {
  return (
    <div className="border border-current-line rounded-lg p-6 space-y-4">
      <div className="flex items-center gap-4">
        {/* Avatar */}
        <div className="w-12 h-12 rounded-full bg-current-line overflow-hidden">
          {author.avatar ? (
            <img
              src={author.avatar}
              alt={author.name}
              className="w-full h-full object-cover"
            />
          ) : (
            <div className="w-full h-full bg-gradient-to-br from-purple/20 to-cyan/20" />
          )}
        </div>

        <div>
          <h3 className="font-semibold">{author.name}</h3>
          <p className="text-sm text-comment">
            {format(date, "d 'de' MMMM, yyyy", { locale: es })}
          </p>
        </div>
      </div>

      <div>
        <h4 className="text-xl font-semibold mb-2">{title}</h4>
        <p className="text-comment">{content}</p>
      </div>

      <div className="flex items-center gap-6 pt-4 border-t border-current-line">
        <button
          onClick={onLike}
          className="flex items-center gap-2 text-comment hover:text-primary transition-colors"
        >
          <Heart className="w-5 h-5" />
          <span>{likes}</span>
        </button>

        <button
          onClick={onComment}
          className="flex items-center gap-2 text-comment hover:text-primary transition-colors"
        >
          <MessageSquare className="w-5 h-5" />
          <span>{comments}</span>
        </button>
      </div>
    </div>
  )
}
