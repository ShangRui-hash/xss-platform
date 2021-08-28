export default{
    GetUserInfo(token){
        let user_info_base64 = token.split(".")[1]
        let user_info_str = window.atob(user_info_base64);
        let user_info = JSON.parse(user_info_str) 
        return user_info
    }
}