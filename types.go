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

type Category struct {
	ID          int64       `json:"category_id"`
	Title       string      `json:"category_title"`
	Description string      `json:"category_description"`
	Links       Links       `json:"links"`
	Permissions Permissions `json:"permissions"`
}

type Links struct {
	Permalink     string `json:"permalink"`
	Detail        string `json:"detail"`
	SubCategories string `json:"sub-categories"`
	SubForums     string `json:"sub-forums"`
}

type Permissions struct {
	View   bool `json:"view"`
	Edit   bool `json:"edit"`
	Delete bool `json:"delete"`
}

type SystemInfo struct {
	VisitorID int64 `json:"visitor_id"`
	Time      int64 `json:"time"`
}

type Errors struct {
	Messages []string `json:"errors"`
}

func (e Errors) Error() string {
	msg := ""
	for _, str := range e.Messages {
		msg = msg + str
	}
	return msg
}
