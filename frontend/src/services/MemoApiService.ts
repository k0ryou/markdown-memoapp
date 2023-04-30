import http from "../common/http-common";

class MemoApiService {
  getAll(md2html: boolean): Promise<any> {
    return http.get(`/api/memos?md2html=${md2html}`);
  }

  get(id: any, md2html: boolean): Promise<any> {
    return http.get(`/api/memos/${id}?md2html=${md2html}`);
  }

  insert(data: any): Promise<any> {
    return http.post("/api/memos", data);
  }

  update(id: any, data: any): Promise<any> {
    return http.put(`/api/memos/${id}`, data);
  }

  delete(id: any): Promise<any> {
    return http.delete(`/api/memos/${id}`);
  }

//   deleteAll(): Promise<any> {
//     return http.delete(`/api/memos`);
//   }

//   findByDescription(title: string): Promise<any> {
//     return http.get(`/api/memos?title=${title}`);
//   }
}

export default new MemoApiService();