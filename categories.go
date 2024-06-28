/* Copyright (C) 2024 d1s
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 */

package lolzapi

import "fmt"

type (
	getCategoriesResp struct {
		Categories      []Category `json:"categories"`
		CategoriesTotal int64      `json:"categories_total"`
		SystemInfo      SystemInfo `json:"system_info"`
	}

	getCategoryResp struct {
		Category   Category   `json:"category"`
		SystemInfo SystemInfo `json:"system_info"`
	}
)

func (c *Client) Categories() ([]Category, error) {
	var cats getCategoriesResp
	return cats.Categories, c.get(nil, &cats, "/categories")
}

func (c *Client) Category(ID int) (Category, error) {
	var cat getCategoryResp
	return cat.Category, c.get(nil, &cat, fmt.Sprintf("/categories/%d", ID))
}
