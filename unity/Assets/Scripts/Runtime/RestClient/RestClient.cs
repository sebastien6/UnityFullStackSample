using System;
using System.Collections;
using UnityEngine;
using UnityEngine.Networking;

namespace RestClient
{
    public class Singleton<T> : MonoBehaviour where T: MonoBehaviour{
        private static T _instance;

        public static T Instance
        {
            get
            {
                if (_instance == null)
                {
                    if (_instance == null)
                    {
                        GameObject obj = new GameObject();
                        _instance = obj.AddComponent<T>();
                        obj.name = typeof(T).Name;
                    }
                }
                return _instance;
            }
        }
    }

   
    public class HttpClient : Singleton<HttpClient>
    {

        /// <summary>
        ///     start a coroutine to Get json payload from an url
        /// </summary>
        /// <param name="url">the url for the get request</param>
        /// <param name="response">Callback action to handle coroutine string response</param>
        public void Get<T>(string url, Action<T> response) where T: class
        {
                StartCoroutine(GetCoroutine(url, response));
        }

        /// <summary>
        ///     IEnumerator used by Get coroutine to execute the webrequest asynchronoussly
        /// </summary>
        /// <param name="url">the url for the get request</param>
        /// <param name="response">Callback action to handle coroutine string response</param>
        private IEnumerator GetCoroutine<T>(string url, Action<T> response) where T: class
        {

            if (Application.internetReachability == NetworkReachability.NotReachable)
            {
                Debug.LogError("Network is not Reachable.");
            }
            
            using(UnityWebRequest webRequest = UnityWebRequestTexture.GetTexture(url))
            {
                
                yield return webRequest.SendWebRequest();
                string contentType = webRequest.GetResponseHeader("Content-Type");

                switch (webRequest.result)
                {
                    case UnityWebRequest.Result.ConnectionError:
                    case UnityWebRequest.Result.DataProcessingError:
                        break;
                    case UnityWebRequest.Result.ProtocolError:
                        Debug.LogError(url + ": HTTP Error: " + webRequest.error);
                        break;
                    case UnityWebRequest.Result.Success:
                        if (typeof(T).Equals(typeof(string))) {
                            if (contentType == "application/json") {
                                response(webRequest.downloadHandler.text as T);
                            } else {
                                Debug.LogError("received data was not in JSON format for URL: " + url);
                            }
                        } else if (typeof(T).Equals(typeof(Texture2D))) {
                            if (contentType == "image/png" || contentType == "image/jpeg") 
                                response(((DownloadHandlerTexture)webRequest.downloadHandler).texture as T);
                        } else {
                            Debug.LogError("received texture was not in JPEG or PNG format for URL: " + url);
                        }

                        break;
                }
            }
        }
    }
}
