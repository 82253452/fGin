import request from '../utils/request';
import {stringify} from 'qs';

export async function query() {
  return request('/api/users');
}
export async function queryCurrent(params) {
  return request(`/sys/currentUser?${stringify(params)}`);
}
export async function queryMenu(params) {
  return request(`/sys/queryMenu?${stringify(params)}`);
}
