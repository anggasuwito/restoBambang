package usecases

import (
	"gomux/main/master/models"
	"gomux/main/master/repositories"
	"gomux/utils"
	"log"
)

//TransaksiUsecaseImpl TransaksiUsecaseImpl
type TransaksiUsecaseImpl struct {
	transaksiRepo repositories.TransaksiRepository
}

//GetAllTransaksi GetTransaksi
func (s TransaksiUsecaseImpl) GetAllTransaksi() ([]*models.Transaksi, error) {
	transaksi, err := s.transaksiRepo.GetAllTransaksi()
	if err != nil {
		return nil, err
	}
	return transaksi, nil
}

//AddTransaksi InsertTransaksi
func (s TransaksiUsecaseImpl) AddTransaksi(newTransaksi models.Transaksi) error {
	err := utils.ValidateInputNotNil(newTransaksi.IDTransaksi, newTransaksi.TanggalTransaksi, newTransaksi.JenisMenu, newTransaksi.NamaMenu, newTransaksi.HargaMenu, newTransaksi.NamaEkstraMenu, newTransaksi.HargaEkstraMenu, newTransaksi.TotalHarga)

	if err != nil {
		return err
	}
	err = s.transaksiRepo.AddTransaksi(newTransaksi)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// //GetTransaksiByID GetTransaksiByID
// func (s TransaksiUsecaseImpl) GetTransaksiByID(id string) (models.Transaksi, error) {
// 	transaksi, err := s.transaksiRepo.GetTransaksiByID(id)
// 	if err != nil {
// 		return transaksi, err
// 	}
// 	return transaksi, nil
// }

// // UpdateTransaksiByID UpdateTransaksiByID
// func (s TransaksiUsecaseImpl) UpdateTransaksiByID(id string, changeTransaksi models.Transaksi) error {
// 	err := s.transaksiRepo.UpdateTransaksiByID(id, changeTransaksi)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //DeleteTransaksiByID DeleteTransaksiByID
// func (s TransaksiUsecaseImpl) DeleteTransaksiByID(id string) error {
// 	err := s.transaksiRepo.DeleteDataTransaksiByID(id)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

//InitTransaksiUseCase InitTransaksiUseCase
func InitTransaksiUseCase(transaksiRepo repositories.TransaksiRepository) TransaksiUseCase {
	return &TransaksiUsecaseImpl{transaksiRepo}
}
