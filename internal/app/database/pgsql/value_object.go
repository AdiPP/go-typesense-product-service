package pgsql

import "fmt"

var findAllProductsByProductIdsQuery = fmt.Sprintf("(%s)", `
select
	rp.product_id,
	rp.product_name,
	rp.product_weight,
	rp.uom_id,
	ru.uom_name,
	ru.uom_abbreviation,
	case
		when rp.is_insurance = 1 then true
		else false
	end as is_insurance,
	case
		when 
		rp.product_preorder_type_id is not null
		and rp.product_preorder_type_id != 0
		and rp.preorder_day != 0
		and rp.preorder_day is not null 
		then true
		else false
	end as is_preorder,
	rp.preorder_day,
	rp.product_preorder_type_id,
	rppt.product_preorder_type_name,
	rp.product_condition_id,
	rpc.product_condition_name,
	rp.product_description,
	rp.store_id,
	rs."name" as store_name,
	rs.slug,
	rs.store_status_id,
	rss.store_status_name as store_status_name,
	rp.date_in + interval '7' hour as upload_date,
	rc.category_id,
	rc.category_code,
	rc.category_name,
	rc.category_slug,
	rp.product_height,
	rp.product_length,
	rp.product_width,
	rp.product_meta_title,
	rp.product_meta_description,
	rp.product_slug,
	rp.status_record
from
	rns_product rp
join rns_uom ru on
	rp.uom_id = ru.uom_id
join rns_product_preorder_type rppt on
	rp.product_preorder_type_id = rppt.product_preorder_type_id
join rns_product_condition rpc on
	rp.product_condition_id = rpc.product_condition_id
join rns_store rs on
	rp.store_id = rs.id
join store.rns_store_status rss on
	rs.store_status_id = rss.store_status_id
join rns_category rc on
	rc.category_id = rp.category_id
where
	ru.status_record <> 'D'
	and rppt.status_record <> 'D'
	and rpc.status_record <> 'D'
	and rss.status_record <> 'D'
	and rc.status_record <> 'D'
`)

var findAllProductSkusByProductIdsQuery = fmt.Sprintf("(%s)", `
select
	case
		when rps.is_default = 1 then true
		else false
	end as is_default,
	rps.is_default,
	rps.product_id,
	rps.product_price,
	rps.product_sku_id,
	rps.product_sku_number,
	rps.product_sku_mpn,
	rps.product_status_id,
	rps2.product_status_name,
	rps.product_weight
from
	rns_product_sku rps
join rns_product_status rps2 on
	rps2.product_status_id = rps.product_status_id
where
	rps.status_record <> 'D'
`)

var findAllProductDiscountsByProductIdsQuery = `
select
	rpd.product_id,
	rps.product_sku_id,
	case
		when rpd.discount_type_id = 1 then rpd.discount_value
		when rps.product_price > 0 then ceil((rpd.discount_value::numeric / rps.product_price) * 100)
	end as "discount_percentage",
	case
		when rpd.discount_type_id = 1 then (rps.product_price -
		ceil((rps.product_price * rpd.discount_value::numeric / 100)))::int4
		else (rps.product_price - rpd.discount_value::numeric)::int4
	end as "discount_price",
	case
		when rpd.discount_type_id = 1 then (rps.product_price * rpd.discount_value::numeric / 100)::int4
		else (rpd.discount_value::numeric)::int4
	end as "discount_amount",
	rpd.discount_type_id,
	rdt.discount_type_name,
	rpd.discount_value
from
	rns_product_discount rpd
join rns_product_sku rps on
	rpd.product_sku_id = rps.product_sku_id
join rns_product rp on
	rps.product_id = rp.product_id
join
	rns_discount_type rdt on
	rdt.discount_type_id = rpd.discount_type_id
where
	rpd.status_record <> 'D'
	and rps.status_record <> 'D'
	and rp.status_record <> 'D'
	and rdt.status_record <> 'D'
	and ((rpd.product_discount_date_start is null
		and rpd.product_discount_date_end is null)
	or (rpd.product_discount_date_start + interval '7 hours' <= now()
		and rpd.product_discount_date_end + interval '7 hours' >= now()))
`
var findAllProductSkuVariantsByProductIdsQuery = fmt.Sprintf("(%s)", `
select
    rpsv.product_id,
	rpsv.product_sku_id,
	rv.variant_name,
	rvv.variant_value
from
	rns_product_sku_variant rpsv
join rns_variant rv on rv.variant_id = rpsv.variant_id
join rns_variant_value rvv on rvv.variant_value_id = rpsv.variant_value_id
join rns_product_sku rps on rpsv.product_sku_id = rps.product_sku_id
join rns_product rp on rps.product_id = rp.product_id
where
    rpsv.status_record <> 'D'
	and rv.status_record <> 'D'
	and rvv.status_record <> 'D'
	and rps.status_record <> 'D'
	and rp.status_record <> 'D'
`)

var findAllProductImagesByProductIdsQuery = fmt.Sprintf("(%s)", `
select 
    rpi.product_id,
    rpi.product_sku_id,
    rpi.image_type_id,
    rpi.product_image,
    rpi.product_image_name,
    rpi.product_image_no_background,
    rpi.product_image_no_background_name
from 
    rns_product_image rpi
join rns_product rp on rpi.product_id = rp.product_id
where rpi.status_record <> 'D'
and rp.status_record <> 'D'
`)

var findAllProductSkuSoldsByProductIdsQuery = `
select
	min(rop.product_id) as product_id,
	rop.product_sku_id,
	sum(rop.order_product_quantity) as total
from
	rns_order_product rop
join rns_order ro on
	rop.order_id = ro.order_id
where
	rop.status_record <> 'D'
	and ro.status_record <> 'D'
	and ro.order_status_id >= 3
	and ro.order_status_id <> 7
group by
	rop.product_sku_id
`

var findAllProductSkuStocksByProductIdsQuery = fmt.Sprintf(
	"(%s)",
	`
select
	min(rps.product_id) as product_id,
	rps.product_sku_id,
	sum(rps.product_stock) as product_stock
from
	rns_product_stock rps
join rns_product rp on
	rps.product_id = rp.product_id
join rns_product_sku rps2 on
	rps.product_sku_id = rps2.product_sku_id
where
	rps.status_record <> 'D'
group by
	rps.product_sku_id
`,
)

var findAllProductSkuReviewsByProductIdsQuery = fmt.Sprintf("(%s)", `
select
	MIN(rrpr.product_id) as product_id,
	rrpr.product_sku_id,
	ROUND(AVG(rrpr.score)) as average_score,
	COUNT(rrpr.id) as reviews_count,
	case
		when ROUND(AVG(rrpr.score)) >= 1
		and ROUND(AVG(rrpr.score)) < 2 then 'Sangat Buruk'
		when ROUND(AVG(rrpr.score)) >= 2
		and ROUND(AVG(rrpr.score)) < 3 then 'Buruk'
		when ROUND(AVG(rrpr.score)) >= 3
		and ROUND(AVG(rrpr.score)) < 4 then 'Cukup Baik'
		when ROUND(AVG(rrpr.score)) >= 4
		and ROUND(AVG(rrpr.score)) < 5 then 'Baik'
		when ROUND(AVG(rrpr.score)) >= 5 then 'Sangat Baik'
		else 'Tidak Ada Skor'
	end as score_message
from
	rns_reviews_product_review rrpr
where
	rrpr.is_deleted = false
group by
	rrpr.product_sku_id
`)

var findAllStoreLocationsByStoreIdsQuery = fmt.Sprintf("(%s)", ` 
select
	rs.store_id,
	rs.city_name
from
	(
	select
		rssl.store_id,
		rssl.city_name,
		row_number() over (partition by rssl.store_id
	order by
		rssl.created_date desc) as rn
	from
		rns_store_store_location rssl
	where
		rssl.is_primary = true
) rs
where
	rs.rn = 1
`)
