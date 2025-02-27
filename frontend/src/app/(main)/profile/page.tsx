'use client'

import { useState, useEffect } from 'react'
import { Button } from '@/components/ui/Button'
import { User, Mail, MapPin, Edit, Settings, Calendar, Heart, Target, Users, Twitter, Instagram, Globe } from 'lucide-react'
import Link from 'next/link'
import { onAuthStateChanged } from 'firebase/auth'
import { auth, db } from '../../../firebase'
import { collection, query, where, getDocs, doc, getDoc } from 'firebase/firestore'
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

interface Node {
  id: string
  title: string
  description: string
  category: string
  current: number
  goal: number
  daysLeft: number
  backers: number
  imageUrl: string
  location: string
}

interface Activity {
  id: string
  type: 'donation' | 'completion'
  amount?: number
  node?: string
  content?: string
  date: string
}

export default function ProfilePage() {
  const [user, setUser] = useState<User | null>(null)
  const [activeTab, setActiveTab] = useState<'nodes' | 'activity' | 'impact'>('nodes')
  const [userNodes, setUserNodes] = useState<Node[]>([])
  const [activities, setActivities] = useState<Activity[]>([])

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, async (currentUser) => {
      if (currentUser) {
        try {
          // Obtener los datos del usuario desde Firestore
          const userDoc = await getDoc(doc(db, 'users', currentUser.uid))
          if (userDoc.exists()) {
            const userData = userDoc.data() as User
            console.log('Datos del usuario cargados:', userData)
            setUser({
              name: userData.name || '',
              email: userData.email || '',
              avatar: userData.avatar || '',
              coverImage: userData.coverImage || '',
              bio: userData.bio || '',
              location: userData.location || '',
              joinDate: userData.joinDate || '',
              links: userData.links || {
                website: '',
                twitter: '',
                instagram: ''
              },
              stats: userData.stats || {
                nodesCreated: 0,
                nodesSupported: 0,
                totalImpact: 0,
                followers: 0
              }
            })
          } else {
            console.log('No se encontraron datos del usuario')
          }

          // Obtener los nodos creados por el usuario
          const nodesQuery = query(collection(db, 'nodes'), where('createdBy', '==', currentUser.uid))
          const querySnapshot = await getDocs(nodesQuery)
          const nodes = querySnapshot.docs.map(doc => ({
            id: doc.id,
            ...doc.data()
          })) as Node[]
          setUserNodes(nodes)

          // Actualizar el contador de nodos creados
          setUser(prevUser => prevUser ? {
            ...prevUser,
            stats: {
              ...prevUser.stats,
              nodesCreated: nodes.length
            }
          } : null)

          // Obtener las actividades del usuario
          const activitiesQuery = query(collection(db, 'activities'), where('userId', '==', currentUser.uid))
          const activitiesSnapshot = await getDocs(activitiesQuery)
          const activities = activitiesSnapshot.docs.map(doc => ({
            id: doc.id,
            ...doc.data()
          })) as Activity[]
          setActivities(activities)
        } catch (error) {
          console.error('Error al obtener los datos del usuario:', error)
        }
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
      <div className="h-64 bg-gradient-to-r from-primary to-secondary rounded-lg overflow-hidden shadow-lg">
        {user.coverImage && (
          <Image
            src={user.coverImage}
            alt="Cover"
            layout="fill"
            objectFit="cover"
            className="opacity-80"
          />
        )}
      </div>

      {/* Profile Header */}
      <div className="relative px-6 pb-6">
        {/* Avatar and Actions */}
        <div className="flex flex-col md:flex-row md:items-end md:justify-between -mt-12 mb-6">
          <div className="flex items-end">
            <div className="w-32 h-32 rounded-full border-4 border-background overflow-hidden bg-gradient-to-br from-purple/20 to-cyan/20 shadow-lg">
              {user.avatar ? (
                <Image
                  src={user.avatar}
                  alt={user.name}
                  width={128}
                  height={128}
                  className="object-cover"
                />
              ) : (
                <div className="w-full h-full flex items-center justify-center">
                  <User className="w-16 h-16 text-comment" />
                </div>
              )}
            </div>
            <div className="ml-4 mb-4">
              <h1 className="text-3xl font-bold text-foreground">{user.name}</h1>
              <p className="text-comment">{user.email}</p>
            </div>
          </div>

          <div className="flex gap-2 mt-4 md:mt-0">
            <Button variant="primary" className="hover:shadow-lg hover:scale-105 transition-transform">
              <Users className="w-4 h-4 mr-2" />
              Seguir
            </Button>
            <Link href="/profile/edit">
              <Button variant="secondary" className="hover:shadow-lg hover:scale-105 transition-transform">
                <Edit className="w-4 h-4 mr-2" />
                Editar Perfil
              </Button>
            </Link>
            <Button variant="ghost" className="hover:shadow-lg hover:scale-105 transition-transform">
              <Settings className="w-4 h-4" />
            </Button>
          </div>
        </div>

        {/* User Info */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          {/* Bio and Details */}
          <div className="md:col-span-2 space-y-6">
            <div className="space-y-4">
              <p className="text-lg text-foreground/90">{user.bio}</p>
              
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
              {[
                { value: user.stats.nodesCreated, label: 'Nodos Creados', color: 'bg-purple/20 text-purple' },
                { value: user.stats.nodesSupported, label: 'Nodos Apoyados', color: 'bg-pink/20 text-pink' },
                { value: `$${user.stats.totalImpact}`, label: 'Impacto Total', color: 'bg-green/20 text-green' },
                { value: user.stats.followers, label: 'Seguidores', color: 'bg-cyan/20 text-cyan' },
              ].map((stat, index) => (
                <div
                  key={index}
                  className={`rounded-lg p-4 text-center shadow-md hover:shadow-lg transition-shadow ${stat.color}`}
                >
                  <p className="text-2xl font-bold">{stat.value}</p>
                  <p className="text-sm text-comment">{stat.label}</p>
                </div>
              ))}
            </div>
          </div>

          {/* Impact Summary */}
          <div className="bg-current-line/20 rounded-lg p-6 shadow-md hover:shadow-lg transition-shadow">
            <h3 className="text-xl font-semibold mb-4 text-foreground">Resumen de Impacto</h3>
            <div className="space-y-4">
              {[
                { icon: <Target className="w-5 h-5 text-primary" />, text: '3 proyectos completados' },
                { icon: <Users className="w-5 h-5 text-cyan" />, text: '250+ personas beneficiadas' },
                { icon: <Heart className="w-5 h-5 text-red" />, text: '89% de satisfacción' },
              ].map((item, index) => (
                <div key={index} className="flex items-center gap-2 text-foreground/90">
                  {item.icon}
                  <span>{item.text}</span>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>

      {/* Tabs */}
      <div className="border-b border-current-line/20 mt-8">
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
                className="group block hover:scale-105 transform transition-transform duration-300"
              >
                <article className="bg-current-line/10 rounded-lg p-6 hover:bg-current-line/20 transition-colors border border-current-line/20 shadow-lg hover:shadow-xl">
                  <div className="flex items-start justify-between mb-4">
                    <div>
                      <span className={`px-3 py-1 text-sm rounded-full ${
                        node.category === 'social'
                          ? 'bg-purple/20 text-purple'
                          : node.category === 'ambiental'
                          ? 'bg-green/20 text-green'
                          : 'bg-pink/20 text-pink'
                      }`}>
                        {node.category}
                      </span>
                      <h3 className="text-2xl font-bold mt-3 group-hover:text-primary transition-colors">
                        {node.title}
                      </h3>
                      <p className="text-foreground/80 mt-2 line-clamp-2">
                        {node.description}
                      </p>
                      <div className="flex items-center gap-2 mt-3 text-foreground/70">
                        <MapPin className="w-5 h-5" />
                        <span>{node.location}</span>
                      </div>
                    </div>
                    {node.imageUrl && (
                      <Image
                        src={node.imageUrl}
                        alt={node.title}
                        width={120}
                        height={120}
                        className="object-cover rounded-lg shadow-md"
                      />
                    )}
                  </div>
                  <div className="flex justify-between items-center mt-6">
                    <div className="text-sm text-foreground/70">
                      <span className="font-semibold">Meta:</span> ${node.goal}
                    </div>
                    <div className="text-sm text-foreground/70">
                      <span className="font-semibold">Días restantes:</span> {node.daysLeft}
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
