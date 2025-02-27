'use client'

import { useState, useEffect } from 'react'
import { Search } from 'lucide-react'
import { categories } from '@/infrastructure/mock/nodeData'
import { db } from '../../../firebase'
import { collection, getDocs, getDoc, doc } from 'firebase/firestore'
import { NodeCard } from '@/components/nodes/NodeCard'

interface Node {
  id: string
  title: string
  category: string
  location: string
  goal: string
  duration: string
  description: string
  rewards: string
  imageUrl: string
  createdAt: Date
  followers: string[]
  createdBy: string
  name: string
  tags: string[]
  createdByName?: string 
}

const getUserNameById = async (userId: string): Promise<string> => {
  try {
    const userDoc = await getDoc(doc(db, 'users', userId))
    return userDoc.exists() ? userDoc.data().name : 'Usuario desconocido'
  } catch (error) {
    console.error('Error fetching user name:', error)
    return 'Usuario desconocido'
  }
}

export default function NodesPage() {
  const [searchQuery, setSearchQuery] = useState('')
  const [selectedCategory, setSelectedCategory] = useState<string | null>(null)
  const [nodes, setNodes] = useState<Node[]>([])
  const [loading, setLoading] = useState(true)

  const fetchNodes = async (category?: string) => {
    try {
      const querySnapshot = await getDocs(collection(db, 'nodes'))
      const nodesData = await Promise.all(querySnapshot.docs
        .filter(doc => !category || doc.data().category === category)
        .map(async doc => {
          const data = doc.data()
          const createdByName = await getUserNameById(data.createdBy) // Obtener el nombre del creador
          return {
            id: doc.id,
            title: data.title || '',
            category: data.category || '',
            location: data.location || '',
            goal: data.goal || '',
            duration: data.duration || '',
            description: data.description || '',
            rewards: data.rewards || '',
            imageUrl: data.imageUrl || '',
            createdAt: data.createdAt?.toDate() || new Date(),
            followers: data.followers || [],
            createdBy: data.createdBy || '',
            name: data.name || '',
            tags: data.tags || [],
            createdByName, // Añadir el nombre del creador al objeto node
          } as Node
        }))
      return nodesData
    } catch (error) {
      console.error('Error fetching nodes:', error)
      return []
    }
  }

  useEffect(() => {
    const fetchNodesData = async () => {
      const nodesData = await fetchNodes(selectedCategory || undefined)
      setNodes(nodesData)
      setLoading(false)
    }
    fetchNodesData()
  }, [selectedCategory])

  if (loading) {
    return <div>Loading...</div>
  }

  return (
    <div className="space-y-8">
      {/* Hero Section */}
      <div className="text-center space-y-4 pb-8 border-b border-current-line">
        <h1 className="text-4xl font-bold">Nodos de Impacto</h1>
        <p className="text-xl text-comment max-w-2xl mx-auto">
          Descubre proyectos sociales y los productos que los apoyan. Conecta con causas que te importan.
        </p>
      </div>

      {/* Filters */}
      <div className="flex flex-col md:flex-row gap-4 items-center">
        <div className="relative flex-1 max-w-xl">
          <input
            type="text"
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            placeholder="Buscar nodos..."
            className="input pl-10 pr-4 py-2 w-full border border-current-line rounded-lg"
          />
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-comment" />
        </div>

        <div className="flex gap-2 overflow-x-auto pb-2 w-full md:w-auto">
          {/* Botón para mostrar todos los nodos (sin filtro) */}
          <button
  onClick={() => setSelectedCategory(null)}
  className={`px-4 py-2 rounded-full text-sm whitespace-nowrap transition-colors ${
    !selectedCategory
      ? 'bg-purple-500 text-white' // Resaltar si no hay categoría seleccionada
      : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
  }`}
>
  Todos
</button>

          {/* Botones de categoría */}
          {categories.map((category) => (
            <button
              key={category.id}
              onClick={() => setSelectedCategory(
                selectedCategory === category.id ? null : category.id
              )}
              className={`px-4 py-2 rounded-full text-sm whitespace-nowrap transition-colors ${
                selectedCategory === category.id
                  ? `border-2 border-${category.color} bg-${category.color}/20 text-${category.color}` // Resaltar la categoría seleccionada
                  : `bg-${category.color}/10 text-${category.color} hover:bg-${category.color}/20`
              }`}
            >
              {category.name}
            </button>
          ))}
        </div>
      </div>

      {/* Nodes Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {nodes
          .filter(node => 
            (selectedCategory ? node.category === selectedCategory : true) &&
            (searchQuery ? node.title.toLowerCase().includes(searchQuery.toLowerCase()) : true)
          )
          .map(node => (
            <NodeCard key={node.id} node={node} />
          ))}
      </div>
    </div>
  )
}