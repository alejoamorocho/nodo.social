import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Search } from "lucide-react";

// Sample product data (replace with actual data fetching)
const products = [
  {
    id: 1,
    title: "Eco-friendly Water Bottle",
    description: "Sustainable stainless steel water bottle with double-wall insulation",
    price: 29.99,
    image: "/images/placeholder.jpg",
    category: "Sustainability",
    stock: 100,
    rating: 4.5,
    reviews: 128,
  },
  // Add more sample products as needed
];

export default function ProductsPage() {
  return (
    <div className="container mx-auto px-4 py-8">
      {/* Hero Section */}
      <div className="mb-12 text-center">
        <h1 className="text-4xl font-bold mb-4">Discover Impact Products</h1>
        <p className="text-gray-600 max-w-2xl mx-auto">
          Browse our curated collection of sustainable and impactful products that make a difference.
        </p>
      </div>

      {/* Search and Filters */}
      <div className="mb-8">
        <div className="flex gap-4 mb-4">
          <div className="relative flex-1">
            <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" />
            <Input
              placeholder="Search products..."
              className="pl-10"
            />
          </div>
          <Button variant="outline">
            Filter
          </Button>
        </div>
        
        {/* Category Pills */}
        <div className="flex gap-2 flex-wrap">
          {["All", "Sustainability", "Education", "Health", "Technology"].map((category) => (
            <Button
              key={category}
              variant="outline"
              className="rounded-full"
              size="sm"
            >
              {category}
            </Button>
          ))}
        </div>
      </div>

      {/* Products Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        {products.map((product) => (
          <div
            key={product.id}
            className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow"
          >
            <div className="aspect-square relative">
              <img
                src={product.image}
                alt={product.title}
                className="object-cover w-full h-full"
              />
            </div>
            <div className="p-4">
              <h3 className="font-semibold text-lg mb-2">{product.title}</h3>
              <p className="text-gray-600 text-sm mb-3 line-clamp-2">
                {product.description}
              </p>
              <div className="flex justify-between items-center">
                <span className="text-xl font-bold">${product.price}</span>
                <div className="flex items-center gap-1">
                  <span className="text-yellow-400">★</span>
                  <span className="text-sm text-gray-600">
                    {product.rating} ({product.reviews})
                  </span>
                </div>
              </div>
              <Button className="w-full mt-4">View Details</Button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
