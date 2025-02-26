'use client'

import { useState, useEffect } from 'react'
import { Search } from 'lucide-react'
import { NodeCard } from '@/components/nodes/NodeCard'
import { categories } from '@/infrastructure/mock/nodeData'
import { db } from '../../../firebase'
import { collection, getDocs, doc, getDoc } from 'firebase/firestore'

interface Product {
  name: string
  price: number
  description: string
}

interface Node {
  id: string
  name: string
  description: string
  category: string
  tags: string[]
  imageUrl: string
  createdAt: Date
  createdBy: string
  products?: Product[] 
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
            className="input pl-10 pr-4 py-2 w-full"
          />
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-comment" />
        </div>

        <div className="flex gap-2 overflow-x-auto pb-2 w-full md:w-auto">
          {categories.map((category) => (
            <button
              key={category.id}
              onClick={() => setSelectedCategory(
                selectedCategory === category.id ? null : category.id
              )}
              className={`px-4 py-2 rounded-full text-sm whitespace-nowrap transition-colors ${
                selectedCategory === category.id
                  ? `bg-${category.color} text-background`
                  : `bg-${category.color}/20 text-${category.color} hover:bg-${category.color}/30`
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
            (searchQuery
              ? node.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
                node.description.toLowerCase().includes(searchQuery.toLowerCase()) ||
                node.tags?.some(tag => tag.toLowerCase().includes(searchQuery.toLowerCase()))
              : true)
          )
          .map(node => (
            <NodeCard key={node.id} node={node} />
          ))
        }
      </div>
    </div>
  )
}