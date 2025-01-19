import { Button } from "@/components/ui/Button";
import { Heart, Share2, Star } from "lucide-react";

// Sample product data (replace with actual data fetching)
const product = {
  id: 1,
  title: "Eco-friendly Water Bottle",
  description: "Our premium stainless steel water bottle is designed with sustainability in mind. Features double-wall vacuum insulation that keeps your drinks cold for 24 hours or hot for 12 hours. Made from 100% recyclable materials, this bottle helps reduce single-use plastic waste.",
  price: 29.99,
  image: "/images/placeholder.jpg",
  category: "Sustainability",
  stock: 100,
  rating: 4.5,
  reviews: 128,
  features: [
    "Double-wall vacuum insulation",
    "18/8 food-grade stainless steel",
    "BPA-free",
    "24-hour cold retention",
    "12-hour hot retention",
    "Leak-proof design"
  ],
  impact: {
    plasticSaved: "500 plastic bottles per year",
    carbonReduction: "62.5 kg CO2 per year",
    waterSaved: "1000 liters per year"
  }
};

export default function ProductPage() {
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        {/* Product Image */}
        <div className="relative aspect-square">
          <img
            src={product.image}
            alt={product.title}
            className="object-cover w-full h-full rounded-lg"
          />
        </div>

        {/* Product Info */}
        <div>
          <div className="flex justify-between items-start">
            <div>
              <h1 className="text-3xl font-bold mb-2">{product.title}</h1>
              <div className="flex items-center gap-2 mb-4">
                <div className="flex items-center">
                  {[...Array(5)].map((_, i) => (
                    <Star
                      key={i}
                      className={`w-4 h-4 ${
                        i < Math.floor(product.rating)
                          ? "text-yellow-400 fill-yellow-400"
                          : "text-gray-300"
                      }`}
                    />
                  ))}
                </div>
                <span className="text-gray-600">
                  {product.rating} ({product.reviews} reviews)
                </span>
              </div>
            </div>
            <div className="flex gap-2">
              <Button variant="outline" size="icon">
                <Heart className="w-4 h-4" />
              </Button>
              <Button variant="outline" size="icon">
                <Share2 className="w-4 h-4" />
              </Button>
            </div>
          </div>

          <div className="text-2xl font-bold mb-6">${product.price}</div>

          <p className="text-gray-600 mb-6">{product.description}</p>

          {/* Features */}
          <div className="mb-6">
            <h2 className="text-xl font-semibold mb-3">Features</h2>
            <ul className="list-disc list-inside space-y-2">
              {product.features.map((feature, index) => (
                <li key={index} className="text-gray-600">{feature}</li>
              ))}
            </ul>
          </div>

          {/* Impact */}
          <div className="mb-6">
            <h2 className="text-xl font-semibold mb-3">Environmental Impact</h2>
            <div className="grid grid-cols-3 gap-4">
              <div className="bg-green-50 p-4 rounded-lg text-center">
                <div className="font-semibold text-green-700">Plastic Saved</div>
                <div className="text-sm text-green-600">{product.impact.plasticSaved}</div>
              </div>
              <div className="bg-green-50 p-4 rounded-lg text-center">
                <div className="font-semibold text-green-700">CO2 Reduction</div>
                <div className="text-sm text-green-600">{product.impact.carbonReduction}</div>
              </div>
              <div className="bg-green-50 p-4 rounded-lg text-center">
                <div className="font-semibold text-green-700">Water Saved</div>
                <div className="text-sm text-green-600">{product.impact.waterSaved}</div>
              </div>
            </div>
          </div>

          {/* Add to Cart */}
          <div className="flex gap-4">
            <Button className="flex-1" size="lg">
              Add to Cart
            </Button>
            <Button variant="outline" size="lg">
              Buy Now
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}
