'use client'

import { format } from 'date-fns'
import { es } from 'date-fns/locale'

interface Backer {
  name: string
  amount: number
  date: Date
  avatar?: string
  comment?: string
}

interface BackersListProps {
  backers: Backer[]
}

export function BackersList({ backers }: BackersListProps) {
  return (
    <div className="space-y-6">
      {backers.map((backer, index) => (
        <div
          key={index}
          className="flex items-start gap-4 p-4 bg-current-line rounded-lg hover:bg-current-line/80 transition-colors"
        >
          {/* Avatar */}
          <div className="w-12 h-12 rounded-full bg-background overflow-hidden flex-shrink-0">
            {backer.avatar ? (
              <img
                src={backer.avatar}
                alt={backer.name}
                className="w-full h-full object-cover"
              />
            ) : (
              <div className="w-full h-full bg-gradient-to-br from-purple/20 to-cyan/20" />
            )}
          </div>

          <div className="flex-1 min-w-0">
            <div className="flex items-start justify-between gap-4">
              <div>
                <h4 className="font-semibold truncate">{backer.name}</h4>
                <p className="text-sm text-comment">
                  {format(backer.date, "d 'de' MMMM, yyyy", { locale: es })}
                </p>
              </div>
              <span className="text-lg font-bold whitespace-nowrap">
                ${backer.amount}
              </span>
            </div>
            
            {backer.comment && (
              <p className="mt-2 text-comment">
                {backer.comment}
              </p>
            )}
          </div>
        </div>
      ))}
    </div>
  )
}
