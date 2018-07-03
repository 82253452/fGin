// use localStorage to store the authority info, which might be sent from server in actual project.
export function getAuthority() {
  return localStorage.getItem('antd-pro-authority');
}

export function getAuthorityRole() {
  return localStorage.getItem('antd-pro-authority') == 'yuping' ? 'admin' : 'user';
}

export function setAuthority(authority) {
  return localStorage.setItem('antd-pro-authority', authority);
}
