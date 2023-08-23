package oracle

import (
	"context"
	"contracts/censor/pkg/config"
	logs "contracts/internal/log"
	"contracts/internal/model"
	"database/sql"

	_ "github.com/godror/godror"

	"fmt"
	"time"
)

const (
	Sql string = `with contr as( select  sd.n_doc_id , sd.n_doc_state_id , sd.d_begin, sd.d_end ,sd.d_log_last_mod ,sd.vc_doc_no, 
		SI_SUBJECTS_PKG_S.GET_BASE_SUBJECT_ID(sds.n_subject_id) as contr , 
		  SI_SUBJECTS_PKG_S.GET_VC_Name(SI_SUBJECTS_PKG_S.GET_BASE_SUBJECT_ID(sds.n_subject_id)) sub
		 -- hd.n_doc_state_id, hd.n_doc_state_id
		  from SD_DOCUMENTS sd , SI_DOC_SUBJECTS sds
		  where   
		 sd.n_doc_type_id = sys_context('CONST','DOC_TYPE_SubscriberContract')
		and sd.n_doc_id = sds.n_doc_id and sds.n_doc_role_id = sys_context('CONST','SUBJ_ROLE_Receiver')
		and    case 
				 when   &1 Is NULL
					then 1
					else 
					  case 
					   when            sd.d_log_last_mod >= TO_DATE (&1, 'DD.MM.YYYY HH24:MI:SS')            
						 then 1
					   else 0
					   end     
			  end = 1                       
			  
		),
		
		history as (SELECT * 
		FROM (SELECT  row_number() over(partition BY hd.n_doc_id  ORDER BY hd.d_log_last_mod ) rn ,
		 hd.n_doc_id,  hd.n_doc_state_id ,hd.d_log_last_mod  
			  FROM   hd_documents hd, contr
			  where hd.n_doc_id = contr.n_doc_id) 
		WHERE rn = 1
			)
		
		select   sd.*, hd.n_doc_state_id last_mod_hist  ,hd.d_log_last_mod date_mod_hist 
		 from contr sd
		left join history  hd on hd.n_doc_id = sd.n_doc_id 
		order by  sd.D_begin

	`
	minDate int64 = 201901010000 //минимальная дата
)

func GetContacts(t int64) (model.Contracts, error) {

	c := config.New()
	param := getParametr(t)
	logs.New().Debug("param ", param)
	constr := fmt.Sprintf("%s/%s@%s:%d/%s", c.DbUser(), c.Passwd(), c.DbHost(), c.DbPort(), c.Db())
	arr := model.New()
	db, err := sql.Open("godror", constr)
	if err != nil {
		logs.New().Error(err)
		return arr, err
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	rows, err := db.QueryContext(ctx, Sql, param)
	if err != nil {
		logs.New().Error(err)
		return arr, err
	}
	defer rows.Close()
	var id, owner, status, statusH *uint64
	var nn, ownerName *string

	var begin, end, dateH, lastmod *time.Time

	for rows.Next() {

		if err = rows.Scan(&id,
			&status,
			&begin,
			&end,
			&lastmod,
			&nn,
			&owner,
			&ownerName,
			&statusH,
			&dateH); err != nil {
			logs.New().Error(err)
			return arr, err
		}
		c := model.CreateContract(id, nn, owner, ownerName, begin, end, status, statusH, dateH)
		arr.Add(*c)
	}

	return arr, rows.Err()
}

func getParametr(p int64) string {
	if p < minDate {
		return ""
	}
	minut := p % 100
	p /= 100
	hour := p % 100
	p /= 100
	day := p % 100
	p /= 100
	mohth := p % 100
	p /= 100
	year := p % 10000
	ret := fmt.Sprintf("%02d.%02d.%04d %02d:%02d:%02d", day, mohth, year, hour, minut, 0) //'12.08.2023 01:01:01'
	return ret
}
