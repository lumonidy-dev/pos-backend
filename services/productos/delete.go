package services

import "log"

func (s *ProductServiceFirestore) DeleteProduct(id string) error {
	_, err := s.client.Collection("Productos").Doc(id).Delete(s.ctx)
	if err != nil {
		log.Printf("Error al eliminar un producto: %v", err)
		return err
	}
	return nil
}
