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
	getForumsResp struct {
		Forums      []Forum    `json:"forums"`
		ForumsTotal int64      `json:"forums_total"`
		SystemInfo  SystemInfo `json:"system_info"`
	}
	getForumResp struct {
		Forum       Forum      `json:"forum"`
		ForumsTotal int64      `json:"forums_total"`
		SystemInfo  SystemInfo `json:"system_info"`
	}
)

func (c *Client) Forums() ([]Forum, error) {
	var forums getForumsResp
	return forums.Forums, c.get(&forums, "/forums")
}

func (c *Client) Forum(ID int) (Forum, error) {
	var forums getForumResp
	return forums.Forum, c.get(&forums, fmt.Sprintf("/forums/%d", ID))
}

func (c *Client) FollowedForums() ([]Forum, error) {
	var forums getForumsResp
	return forums.Forums, c.get(&forums, "/forums/followed")
}

func (c *Client) FollowForum(ID int) error {
	return c.post(nil, nil, fmt.Sprintf("/forums/%d/followers", ID))
}

func (c *Client) UnfollowForum(ID int) error {
	return c.post(nil, nil, fmt.Sprintf("/forums/%d/followers", ID))
}
