package services

import "log"

func (s *CategoryServiceFirestore) DeleteCategory(id string) error {
	_, err := s.client.Collection("Categorias").Doc(id).Delete(s.ctx)
	if err != nil {
		log.Printf("Error al eliminar una categoria: %v", err)
		return err
	}
	return nil
}
