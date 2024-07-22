package server

func (s *Server) Cancel() error {
	if err := s.db.Close(); err != nil {
		return err
	}

	return nil
}
