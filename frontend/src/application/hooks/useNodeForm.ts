import { useState } from 'react'

export interface NodeFormData {
  title: string
  category: string
  location: string
  goal: string
  duration: string
  description: string
  rewards: {
    title: string
    amount: string
    description: string
    items: string[]
    quantity?: string
    estimatedDelivery?: string
  }[]
}

const initialFormData: NodeFormData = {
  title: '',
  category: '',
  location: '',
  goal: '',
  duration: '',
  description: '',
  rewards: [
    {
      title: '',
      amount: '',
      description: '',
      items: ['']
    }
  ]
}

export const useNodeForm = () => {
  const [formData, setFormData] = useState<NodeFormData>(initialFormData)
  const [currentStep, setCurrentStep] = useState(1)

  const handleInputChange = (field: keyof NodeFormData, value: string) => {
    setFormData(prev => ({ ...prev, [field]: value }))
  }

  const handleRewardChange = (index: number, field: string, value: string) => {
    setFormData(prev => ({
      ...prev,
      rewards: prev.rewards.map((reward, i) =>
        i === index ? { ...reward, [field]: value } : reward
      )
    }))
  }

  const addReward = () => {
    setFormData(prev => ({
      ...prev,
      rewards: [
        ...prev.rewards,
        {
          title: '',
          amount: '',
          description: '',
          items: ['']
        }
      ]
    }))
  }

  const removeReward = (index: number) => {
    setFormData(prev => ({
      ...prev,
      rewards: prev.rewards.filter((_, i) => i !== index)
    }))
  }

  const nextStep = () => setCurrentStep(prev => prev + 1)
  const prevStep = () => setCurrentStep(prev => prev - 1)

  return {
    formData,
    currentStep,
    handleInputChange,
    handleRewardChange,
    addReward,
    removeReward,
    nextStep,
    prevStep
  }
}
