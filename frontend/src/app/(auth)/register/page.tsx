import { Card } from '@/components/ui/Card'
import { Button } from '@/components/ui/Button'

export default function RegisterPage() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-background p-4">
      <Card className="w-full max-w-md">
        <h1 className="text-2xl font-bold text-foreground mb-6">Crear Cuenta</h1>
        <form className="space-y-4">
          <div>
            <label htmlFor="name" className="block text-sm font-medium text-foreground mb-1">
              Nombre
            </label>
            <input
              type="text"
              id="name"
              className="input w-full"
              placeholder="Tu nombre"
            />
          </div>
          <div>
            <label htmlFor="email" className="block text-sm font-medium text-foreground mb-1">
              Email
            </label>
            <input
              type="email"
              id="email"
              className="input w-full"
              placeholder="tu@email.com"
            />
          </div>
          <div>
            <label htmlFor="password" className="block text-sm font-medium text-foreground mb-1">
              Contraseña
            </label>
            <input
              type="password"
              id="password"
              className="input w-full"
              placeholder="••••••••"
            />
          </div>
          <Button className="w-full">
            Crear Cuenta
          </Button>
        </form>
      </Card>
    </div>
  )
}
