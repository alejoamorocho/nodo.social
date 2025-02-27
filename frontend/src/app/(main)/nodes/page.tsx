'use client'

import { useState, useEffect } from 'react'
import { Search, X } from 'lucide-react'
import { categories } from '@/infrastructure/mock/nodeData'
import { db } from '../../../firebase'
import { collection, getDocs, doc, getDoc } from 'firebase/firestore'
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
}

interface User {
  id: string
  name: string
}

export default function NodesPage() {
  const [searchQuery, setSearchQuery] = useState('')
  const [selectedCategory, setSelectedCategory] = useState<string | null>(null)
  const [nodes, setNodes] = useState<Node[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchNodes = async () => {
      try {
        const querySnapshot = await getDocs(collection(db, 'nodes'))
        const nodesData = await Promise.all(querySnapshot.docs.map(async (nodeDoc) => {
          const nodeData = nodeDoc.data() as Node
          const userDoc = await getDoc(doc(db, 'users', nodeData.createdBy))
          const userData = userDoc.exists() ? userDoc.data() as User : { name: 'Usuario desconocido' }
          return { ...nodeData, createdBy: userData.name }
        }))
        setNodes(nodesData)
      } catch (error) {
        console.error('Error fetching nodes:', error)
      } finally {
        setLoading(false)
      }
    }

    fetchNodes()
  }, [])

  // Función para limpiar los filtros
  const clearFilters = () => {
    setSelectedCategory(null)
    setSearchQuery('')
  }

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
          {/* Botón para limpiar filtros */}
          {selectedCategory && (
            <button
              onClick={clearFilters}
              className="flex items-center gap-2 px-4 py-2 rounded-full text-sm bg-gray-100 text-gray-700 hover:bg-gray-200 transition-colors border border-gray-300"
            >
              <X className="w-4 h-4" />
              Quitar filtros
            </button>
          )}

          {/* Botones de categoría */}
          {categories.map((category) => (
            <button
              key={category.id}
              onClick={() => setSelectedCategory(
                selectedCategory === category.id ? null : category.id
              )}
              className={`px-4 py-2 rounded-full text-sm whitespace-nowrap transition-colors ${
                selectedCategory === category.id
                  ? `border-2 border-${category.color} bg-${category.color}/20 text-${category.color}`
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