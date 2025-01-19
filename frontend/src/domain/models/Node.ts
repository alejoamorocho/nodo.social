export interface Node {
  id: string
  name: string
  description: string
  category: string
  location: string
  imageUrl: string
  socialMetrics: {
    connections: number
    shares: number
    points: number
    comments: number
  }
  products: Product[]
  contacts: {
    instagram?: string
    whatsapp?: string
    email: string
  }
  tags: string[]
}

export interface Product {
  id: string
  title: string
  description: string
  imageUrl: string
  nodeContribution: number
  contacts: {
    instagram?: string
    whatsapp?: string
  }
  tags: string[]
}

export interface UserGamification {
  points: number
  level: number
  rewards: Reward[]
  achievements: Achievement[]
}

export interface Reward {
  id: string
  title: string
  description: string
  points: number
  achieved: boolean
}

export interface Achievement {
  id: string
  title: string
  description: string
  progress: number
  total: number
  achieved: boolean
}
