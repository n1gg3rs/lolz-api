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

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const Host = "api.zelenka.guru"

type Client struct {
	token string
	host  string
	c     *http.Client
}

func New(token string) *Client {
	return &Client{
		token: token,
		host:  Host,
		c:     &http.Client{},
	}
}

func (c *Client) get(v, dest any, end string) error {
	uri := fmt.Sprintf("https://%s%s", c.host, end)
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer "+c.token)

	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == 403 {
		var e Errors
		err := json.Unmarshal(body, &e)
		if err != nil {
			return err
		}
		return e
	}

	return json.Unmarshal(body, dest)
}
