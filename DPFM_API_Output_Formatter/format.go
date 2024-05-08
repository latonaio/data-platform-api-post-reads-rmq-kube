package dpfm_api_output_formatter

import (
	"data-platform-api-post-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.Post,
			&pm.PostType,
			&pm.PostOwner,
			&pm.PostOwnerBusinessPartnerRole,
			&pm.Description,
			&pm.LongText,
			&pm.Tag1,
			&pm.Tag2,
			&pm.Tag3,
			&pm.Tag4,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsMarkedForDeletion,

		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			Post:							data.Post,
			PostType:						data.PostType,
			PostOwner:						data.PostOwner,
			PostOwnerBusinessPartnerRole:	data.PostOwnerBusinessPartnerRole,
			Description:					data.Description,
			LongText:						data.LongText,
			Tag1:							data.Tag1,
			Tag2:							data.Tag2,
			Tag3:							data.Tag3,
			Tag4:							data.Tag4,
			CreationDate:					data.CreationDate,
			CreationTime:					data.CreationTime,
			LastChangeDate:					data.LastChangeDate,
			LastChangeTime:					data.LastChangeTime,
			IsMarkedForDeletion:			data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToFriend(rows *sql.Rows) (*[]Friend, error) {
	defer rows.Close()
	friend := make([]Friend, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Friend{}

		err := rows.Scan(
			&pm.Post,
			&pm.Friend,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &friend, err
		}

		data := pm
		friend = append(friend, Friend{
			Post:					data.Post,
			Friend:					data.Friend,
			CreationDate:			data.CreationDate,
			CreationTime:			data.CreationTime,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &friend, nil
	}

	return &friend, nil
}
