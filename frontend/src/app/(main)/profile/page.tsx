'use client'

import { useState, useEffect } from 'react'
import { Button } from '@/components/ui/Button'
import { User, Mail, MapPin, Edit, Settings, Calendar, Heart, Target, Users, Twitter, Instagram, Globe } from 'lucide-react'
import Link from 'next/link'
import { onAuthStateChanged } from 'firebase/auth'
import { auth } from '../../../../firebase'
import Image from 'next/image'

interface User {
  name: string
  email: string
  avatar: string
  coverImage: string
  bio: string
  location: string
  joinDate: string
  links: {
    website: string
    twitter: string
    instagram: string
  }
  stats: {
    nodesCreated: number
    nodesSupported: number
    totalImpact: number
    followers: number
  }
}

const userNodes = [
  {
    id: 1,
    type: 'created',
    title: 'Refugio Animal Sustentable',
    description: 'Un espacio seguro para animales rescatados',
    category: 'animal',
    current: 15000,
    goal: 50000,
    daysLeft: 45,
    backers: 128
  },
  {
    id: 2,
    type: 'supported',
    title: 'Huertos Urbanos Comunitarios',
    description: 'Transformando espacios urbanos en jardines productivos',
    category: 'ambiental',
    current: 8000,
    goal: 20000,
    daysLeft: 30,
    backers: 89
  }
]

const activities = [
  {
    id: 1,
    type: 'donation',
    node: 'Huertos Urbanos Comunitarios',
    amount: 500,
    date: '2 días atrás'
  },
  {
    id: 2,
    type: 'update',
    node: 'Refugio Animal Sustentable',
    content: 'Publicó una actualización: ¡Alcanzamos el 30%!',
    date: '5 días atrás'
  }
]

export default function ProfilePage() {
  const [user, setUser] = useState<User | null>(null)
  const [activeTab, setActiveTab] = useState<'nodes' | 'activity' | 'impact'>('nodes')

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, (currentUser) => {
      if (currentUser) {
        setUser({
          name: currentUser.displayName || '',
          email: currentUser.email || '',
          avatar: currentUser.photoURL || '',
          coverImage: '', // Agrega la propiedad coverImage si es necesario
          bio: '', // Agrega la propiedad bio si es necesario
          location: '', // Agrega la propiedad location si es necesario
          joinDate: '', // Agrega la propiedad joinDate si es necesario
          links: {
            website: '',
            twitter: '',
            instagram: ''
          },
          stats: {
            nodesCreated: 0,
            nodesSupported: 0,
            totalImpact: 0,
            followers: 0
          }
        })
      } else {
        setUser(null)
      }
    })

    return () => unsubscribe()
  }, [])

  if (!user) {
    return <div>Loading...</div>
  }

  return (
    <div className="max-w-6xl mx-auto">
      {/* Cover Image */}
      <div className="h-64 bg-gradient-to-r from-primary to-secondary rounded-lg overflow-hidden">
        {user.coverImage && (
          <Image
            src={user.coverImage}
            alt="Cover"
            layout="fill"
            objectFit="cover"
          />
        )}
      </div>

      {/* Profile Header */}
      <div className="relative px-6 pb-6">
        {/* Avatar and Actions */}
        <div className="flex flex-col md:flex-row md:items-end md:justify-between -mt-12 mb-6">
          <div className="flex items-end">
            <div className="w-32 h-32 rounded-full border-4 border-background overflow-hidden bg-current-line">
              {user.avatar ? (
                <Image
                  src={user.avatar}
                  alt={user.name}
                  width={128}
                  height={128}
                  className="object-cover"
                />
              ) : (
                <div className="w-full h-full bg-gradient-to-br from-purple/20 to-cyan/20 flex items-center justify-center">
                  <User className="w-16 h-16 text-comment" />
                </div>
              )}
            </div>
            <div className="ml-4 mb-4">
              <h1 className="text-3xl font-bold">{user.name}</h1>
              <p className="text-comment">{user.email}</p>
            </div>
          </div>

          <div className="flex gap-2 mt-4 md:mt-0">
            <Button variant="primary">
              <Users className="w-4 h-4 mr-2" />
              Seguir
            </Button>
            <Button variant="secondary">
              <Edit className="w-4 h-4 mr-2" />
              Editar Perfil
            </Button>
            <Button variant="ghost">
              <Settings className="w-4 h-4" />
            </Button>
          </div>
        </div>

        {/* User Info */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          {/* Bio and Details */}
          <div className="md:col-span-2 space-y-6">
            <div className="space-y-4">
              <p className="text-lg">{user.bio}</p>
              
              <div className="flex flex-wrap gap-4 text-comment">
                <span className="flex items-center gap-1">
                  <MapPin className="w-4 h-4" />
                  {user.location}
                </span>
                <span className="flex items-center gap-1">
                  <Mail className="w-4 h-4" />
                  {user.email}
                </span>
                <span className="flex items-center gap-1">
                  <Calendar className="w-4 h-4" />
                  Se unió en {user.joinDate}
                </span>
              </div>

              <div className="flex flex-wrap gap-4">
                {user.links.website && (
                  <a href={user.links.website} className="flex items-center gap-1 text-comment hover:text-primary transition-colors">
                    <Globe className="w-4 h-4" />
                    Sitio web
                  </a>
                )}
                {user.links.twitter && (
                  <a href={`https://twitter.com/${user.links.twitter}`} className="flex items-center gap-1 text-comment hover:text-primary transition-colors">
                    <Twitter className="w-4 h-4" />
                    {user.links.twitter}
                  </a>
                )}
                {user.links.instagram && (
                  <a href={`https://instagram.com/${user.links.instagram}`} className="flex items-center gap-1 text-comment hover:text-primary transition-colors">
                    <Instagram className="w-4 h-4" />
                    {user.links.instagram}
                  </a>
                )}
              </div>
            </div>

            {/* Stats */}
            <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
              <div className="bg-current-line rounded-lg p-4 text-center">
                <p className="text-2xl font-bold">{user.stats.nodesCreated}</p>
                <p className="text-sm text-comment">Nodos Creados</p>
              </div>
              <div className="bg-current-line rounded-lg p-4 text-center">
                <p className="text-2xl font-bold">{user.stats.nodesSupported}</p>
                <p className="text-sm text-comment">Nodos Apoyados</p>
              </div>
              <div className="bg-current-line rounded-lg p-4 text-center">
                <p className="text-2xl font-bold">${user.stats.totalImpact}</p>
                <p className="text-sm text-comment">Impacto Total</p>
              </div>
              <div className="bg-current-line rounded-lg p-4 text-center">
                <p className="text-2xl font-bold">{user.stats.followers}</p>
                <p className="text-sm text-comment">Seguidores</p>
              </div>
            </div>
          </div>

          {/* Impact Summary */}
          <div className="bg-current-line rounded-lg p-6">
            <h3 className="text-xl font-semibold mb-4">Resumen de Impacto</h3>
            <div className="space-y-4">
              <div className="flex items-center gap-2">
                <Target className="w-5 h-5 text-primary" />
                <span>3 proyectos completados</span>
              </div>
              <div className="flex items-center gap-2">
                <Users className="w-5 h-5 text-cyan" />
                <span>250+ personas beneficiadas</span>
              </div>
              <div className="flex items-center gap-2">
                <Heart className="w-5 h-5 text-red" />
                <span>89% de satisfacción</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Tabs */}
      <div className="border-b border-current-line mt-8">
        <nav className="flex gap-8">
          {[
            { id: 'nodes', label: 'Nodos' },
            { id: 'activity', label: 'Actividad' },
            { id: 'impact', label: 'Impacto' }
          ].map((tab) => (
            <button
              key={tab.id}
              onClick={() => setActiveTab(tab.id as typeof activeTab)}
              className={`pb-4 text-comment hover:text-foreground transition-colors relative ${
                activeTab === tab.id ? 'text-foreground after:absolute after:bottom-0 after:left-0 after:w-full after:h-0.5 after:bg-primary' : ''
              }`}
            >
              {tab.label}
            </button>
          ))}
        </nav>
      </div>

      {/* Tab Content */}
      <div className="py-8">
        {activeTab === 'nodes' && (
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {userNodes.map((node) => (
              <Link
                key={node.id}
                href={`/nodes/${node.id}`}
                className="group"
              >
                <article className="bg-current-line rounded-lg p-6 hover:bg-current-line/80 transition-colors">
                  <div className="flex items-start justify-between mb-4">
                    <div>
                      <span className={`px-2 py-1 text-xs rounded-full ${
                        node.type === 'created' ? 'bg-primary/20 text-primary' : 'bg-secondary/20 text-secondary'
                      }`}>
                        {node.type === 'created' ? 'Creado' : 'Apoyado'}
                      </span>
                      <h3 className="text-xl font-semibold mt-2 group-hover:text-primary transition-colors">
                        {node.title}
                      </h3>
                      <p className="text-comment mt-1">
                        {node.description}
                      </p>
                    </div>
                  </div>
                </article>
              </Link>
            ))}
          </div>
        )}

        {activeTab === 'activity' && (
          <div className="space-y-6">
            {activities.map((activity) => (
              <div
                key={activity.id}
                className="bg-current-line rounded-lg p-6 flex items-start gap-4"
              >
                <div className="w-10 h-10 rounded-full bg-primary/20 flex items-center justify-center flex-shrink-0">
                  {activity.type === 'donation' ? (
                    <Heart className="w-5 h-5 text-primary" />
                  ) : (
                    <Target className="w-5 h-5 text-primary" />
                  )}
                </div>
                <div className="flex-1">
                  <p>
                    {activity.type === 'donation' ? (
                      <>
                        Donó <span className="font-semibold">${activity.amount}</span> a{' '}
                        <Link href="#" className="text-primary hover:underline">
                          {activity.node}
                        </Link>
                      </>
                    ) : (
                      activity.content
                    )}
                  </p>
                  <span className="text-sm text-comment">{activity.date}</span>
                </div>
              </div>
            ))}
          </div>
        )}

        {activeTab === 'impact' && (
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div className="bg-current-line rounded-lg p-6">
              <h3 className="text-xl font-semibold mb-4">Impacto Social</h3>
              <div className="space-y-4">
                <div className="flex justify-between items-center">
                  <span>Personas beneficiadas</span>
                  <span className="font-semibold">250+</span>
                </div>
                <div className="flex justify-between items-center">
                  <span>Comunidades alcanzadas</span>
                  <span className="font-semibold">5</span>
                </div>
                <div className="flex justify-between items-center">
                  <span>Voluntarios conectados</span>
                  <span className="font-semibold">45</span>
                </div>
              </div>
            </div>

            <div className="bg-current-line rounded-lg p-6">
              <h3 className="text-xl font-semibold mb-4">Impacto Ambiental</h3>
              <div className="space-y-4">
                <div className="flex justify-between items-center">
                  <span>Árboles plantados</span>
                  <span className="font-semibold">100</span>
                </div>
                <div className="flex justify-between items-center">
                  <span>kg de residuos reciclados</span>
                  <span className="font-semibold">500</span>
                </div>
                <div className="flex justify-between items-center">
                  <span>Huertos comunitarios</span>
                  <span className="font-semibold">2</span>
                </div>
              </div>
            </div>

            <div className="bg-current-line rounded-lg p-6">
              <h3 className="text-xl font-semibold mb-4">Impacto Animal</h3>
              <div className="space-y-4">
                <div className="flex justify-between items-center">
                  <span>Animales rescatados</span>
                  <span className="font-semibold">75</span>
                </div>
                <div className="flex justify-between items-center">
                  <span>Adopciones facilitadas</span>
                  <span className="font-semibold">45</span>
                </div>
                <div className="flex justify-between items-center">
                  <span>Campañas de esterilización</span>
                  <span className="font-semibold">3</span>
                </div>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  )
}
