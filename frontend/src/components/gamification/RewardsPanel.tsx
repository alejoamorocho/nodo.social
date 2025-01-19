'use client'

import { Trophy, Star, Award } from 'lucide-react'

interface Reward {
  title: string
  description: string
  points: number
  achieved: boolean
}

interface RewardsPanelProps {
  currentPoints: number
  level: number
  rewards: Reward[]
}

export function RewardsPanel({ currentPoints, level, rewards }: RewardsPanelProps) {
  return (
    <div className="space-y-6">
      {/* User Level */}
      <div className="flex items-center gap-4 p-4 bg-current-line rounded-lg">
        <Trophy className="w-8 h-8 text-yellow-500" />
        <div>
          <p className="text-xl font-bold">Nivel {level}</p>
          <p className="text-comment">{currentPoints.toLocaleString()} puntos</p>
        </div>
      </div>

      {/* Rewards List */}
      <div className="space-y-4">
        <h3 className="text-lg font-bold flex items-center gap-2">
          <Award className="w-5 h-5" />
          Recompensas Disponibles
        </h3>

        <div className="space-y-3">
          {rewards.map((reward, index) => (
            <div
              key={index}
              className={`p-4 rounded-lg border-2 ${
                reward.achieved
                  ? 'border-green-500 bg-green-500/10'
                  : 'border-current-line'
              }`}
            >
              <div className="flex justify-between items-start">
                <div>
                  <h4 className="font-bold">{reward.title}</h4>
                  <p className="text-sm text-comment">{reward.description}</p>
                </div>
                <div className="flex items-center gap-1 text-yellow-500">
                  <Star className="w-4 h-4" />
                  <span>{reward.points}</span>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}
