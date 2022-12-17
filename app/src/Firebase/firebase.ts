import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

// const firebaseConfig = {
//   apiKey: process.env.REACT_APP_API_KEY,
//   authDomain: process.env.REACT_APP_AUTH_DOMAIN,
//   projectId: process.env.REACT_APP_PROJECT_ID,
//   storageBucket: process.env.REACT_APP_STORAGE_BUCKET,
//   messagingSenderId: process.env.REACT_APP_MESSAGEING_SENDER_ID,
//   appId: process.env.REACT_APP_APP_ID,
// };

////上記の.envファイル読み込みはうまくいかず下記なら大丈夫だった
const firebaseConfig = {
  apiKey: "AIzaSyBf50KsuyfJqC1roQm7zvk-sV-tG4TdtI0",
  authDomain: "subtle-melody-368910.firebaseapp.com",
  projectId: "subtle-melody-368910",
  storageBucket: "subtle-melody-368910.appspot.com",
  messagingSenderId: "2213730883",
  appId: "1:2213730883:web:de9b792c3aaf28e63612f6"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);

export const fireAuth = getAuth(app);
