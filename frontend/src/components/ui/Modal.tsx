import React from 'react';
import { Button } from './Button';

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  title: string;
  children: React.ReactNode;
}

const Modal: React.FC<ModalProps> = ({ isOpen, onClose, title, children }) => {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div className="bg-background rounded-lg shadow-lg p-6 w-full max-w-md border border-current-line animate-fade-in">
        {/* Encabezado del modal */}
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-bold text-foreground">{title}</h2>
          <button
            onClick={onClose}
            className="text-foreground/50 hover:text-foreground transition-colors"
          >
            &times;
          </button>
        </div>

        {/* Contenido del modal */}
        <div className="text-foreground/80">{children}</div>

        {/* Pie del modal */}
        <div className="mt-6 flex justify-end">
        <Button
            onClick={onClose}
            className="btn btn-primary"
          >
            Cerrar
          </Button>
        </div>
      </div>
    </div>
  );
};

export default Modal;