import { Node } from '@/domain/models/Node'

export const categories = [
  { id: 'social', name: 'Social', color: 'purple' },
  { id: 'ambiental', name: 'Ambiental', color: 'green' },
  { id: 'animal', name: 'Animal', color: 'pink' },
  { id: 'educacion', name: 'Educación', color: 'cyan' },
  { id: 'salud', name: 'Salud', color: 'yellow' }
]

export const mockNodes: Node[] = [
  {
    id: '1',
    name: 'Refugio Animal Sustentable',
    description: 'Creando un espacio seguro y autosustentable para animales rescatados',
    category: 'animal',
    location: 'Ciudad de México',
    imageUrl: '/placeholder-1.jpg',
    socialMetrics: {
      connections: 128,
      shares: 45,
      points: 350,
      comments: 67
    },
    products: [
      {
        id: 'p1',
        title: 'Camiseta Solidaria',
        description: 'Camiseta 100% algodón orgánico con diseños de animales rescatados',
        imageUrl: '/products/shirt-1.jpg',
        nodeContribution: 30,
        contacts: {
          instagram: 'eco_shirts',
          whatsapp: '5215512345678'
        },
        tags: ['ropa', 'sustentable', 'mascotas']
      },
      {
        id: 'p2',
        title: 'Collar Artesanal',
        description: 'Collares hechos a mano con materiales reciclados',
        imageUrl: '/products/collar-1.jpg',
        nodeContribution: 25,
        contacts: {
          instagram: 'pet_crafts',
          whatsapp: '5215587654321'
        },
        tags: ['mascotas', 'artesanal', 'reciclado']
      }
    ],
    contacts: {
      instagram: 'refugio_sustentable',
      whatsapp: '5215523456789',
      email: 'contacto@refugiosustentable.org'
    },
    tags: ['animales', 'sustentabilidad', 'rescate']
  },
  {
    id: '2',
    name: 'Huertos Urbanos Comunitarios',
    description: 'Transformando espacios urbanos en jardines productivos para la comunidad',
    category: 'ambiental',
    location: 'Guadalajara',
    imageUrl: '/placeholder-2.jpg',
    socialMetrics: {
      connections: 89,
      shares: 32,
      points: 280,
      comments: 45
    },
    products: [
      {
        id: 'p3',
        title: 'Kit de Cultivo Urbano',
        description: 'Todo lo necesario para comenzar tu propio huerto en casa',
        imageUrl: '/products/garden-kit.jpg',
        nodeContribution: 20,
        contacts: {
          instagram: 'urban_gardens',
          whatsapp: '5215534567890'
        },
        tags: ['jardín', 'sustentable', 'cultivo']
      }
    ],
    contacts: {
      instagram: 'huertos_urbanos',
      whatsapp: '5215545678901',
      email: 'info@huertosurbanos.org'
    },
    tags: ['agricultura', 'comunidad', 'sustentabilidad']
  }
]
