import axios from "axios";
import { v4 as uuid } from 'uuid';


export const uploadAudio = async (blob: Blob, displayName: string) => {
    const formData = new FormData();
    const fileName = `${uuid()}.wav`;
    formData.append("file", blob, fileName);
    formData.append("display_name", displayName);
    const response = await axios.post('http://localhost:8080/api/v1/models/inference', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
    return response.data;
    };


export const getModels = async () => {
      const response = await axios.get('http://localhost:8080/api/v1/models/', {headers: {
        'Content-Type': 'application/json',
      }});
      return response.data;
  }
  