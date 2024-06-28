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

type Forum struct {
	ForumID                int64         `json:"forum_id"`
	ForumTitle             string        `json:"forum_title"`
	ForumDescription       string        `json:"forum_description"`
	ForumThreadCount       int64         `json:"forum_thread_count"`
	ForumPostCount         int64         `json:"forum_post_count"`
	ForumPrefixes          []ForumPrefix `json:"forum_prefixes"`
	ThreadDefaultPrefixID  int64         `json:"thread_default_prefix_id"`
	ThreadPrefixIsRequired bool          `json:"thread_prefix_is_required"`
	Links                  Links         `json:"links"`
	Permissions            Permissions   `json:"permissions"`
	ForumIsFollowed        bool          `json:"forum_is_followed"`
}

type ForumPrefix struct {
	GroupTitle    string        `json:"group_title"`
	GroupPrefixes []GroupPrefix `json:"group_prefixes"`
}

type GroupPrefix struct {
	PrefixID    int64  `json:"prefix_id"`
	PrefixTitle string `json:"prefix_title"`
}

type Links struct {
	Permalink     string `json:"permalink"`
	Detail        string `json:"detail"`
	SubCategories string `json:"sub-categories"`
	SubForums     string `json:"sub-forums"`
	Threads       string `json:"threads",omitempty`
	Followers     string `json:"followers",omitempty`
}

type Permissions struct {
	View             bool `json:"view"`
	Edit             bool `json:"edit"`
	Delete           bool `json:"delete"`
	CreateThread     bool `json:"create_thread",omitempty`
	UploadAttachment bool `json:"upload_attachment",omitempty`
	TagThread        bool `json:"tag_thread",omitempty`
	Follow           bool `json:"follow",omitempty`
}

type Category struct {
	ID          int64       `json:"category_id"`
	Title       string      `json:"category_title"`
	Description string      `json:"category_description"`
	Links       Links       `json:"links"`
	Permissions Permissions `json:"permissions"`
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
