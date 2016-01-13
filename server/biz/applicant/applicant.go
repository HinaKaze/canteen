package applicant

import (
	"HKCanteen/server/dao"
)

type Applicant struct {
	Applicant dao.DAOApplicant
	User      dao.DAOUser
}

func JoinOrder(userId int64, orderId int64) (result int) {
	//	var order dao.DAOOrder
	//	order.FetchFromDB(orderId)
	//	if order.Id <= 0 {
	//		return 0
	//	}
	applicant := dao.GetApplicantByOrderId(userId, orderId)
	if applicant.Id > 0 {
		if applicant.Status != "join" {
			applicant.Status = "join"
			applicant.UpdateToDB()
		}
	} else {
		applicant.OrderId = orderId
		applicant.Status = "join"
		applicant.UserId = userId
		applicant.SaveToDB()
	}
	return 1
}

func QuitOrder(userId int64, orderId int64) (result int) {
	//	var order dao.DAOOrder
	//	order.FetchFromDB(orderId)
	//	if order.Id <= 0 {
	//		return 0
	//	}
	applicant := dao.GetApplicantByOrderId(userId, orderId)
	if applicant.Id > 0 {
		if applicant.Status != "quit" {
			applicant.Status = "quit"
			applicant.UpdateToDB()
		}
	} else {
		applicant.OrderId = orderId
		applicant.Status = "quit"
		applicant.UserId = userId
		applicant.SaveToDB()
	}
	return 1
}

func GetApplicantsByOrderId(orderId int64) (applicants []Applicant) {
	daoApplicants := dao.GetApplicantsByOrderId(orderId)
	for i := range daoApplicants {
		var a Applicant
		a.User.FetchFromDB(daoApplicants[i].UserId)
		a.Applicant = daoApplicants[i]
		applicants = append(applicants, a)
	}
	return
}