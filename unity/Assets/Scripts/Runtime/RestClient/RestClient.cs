using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Networking;

namespace RestClient
{
    public class HttpClient : MonoBehaviour
    {
        private static HttpClient _instance;
        public static HttpClient Instance
        {
            get
            {
                if (_instance == null)
                {
                    if (_instance == null)
                    {
                        GameObject obj = new GameObject();
                        obj.name = string.Format("_{0}", typeof(HttpClient).Name);
                        _instance = obj.AddComponent<HttpClient>();
                    }
                }
                return _instance;
            }
        }

        /// <summary>
        ///     start a coroutine to Get json payload from an url
        /// </summary>
        /// <param name="url">the url for the get request</param>
        /// <param name="response">Callback action to handle coroutine string response</param>
        public void Get(string url, Action<string> response)
        {
            StartCoroutine(GetCoroutine(url, response));
        }

        /// <summary>
        ///     IEnumerator used by Get coroutine to execute the webrequest asynchronously
        /// </summary>
        /// <param name="url">the url for the get request</param>
        /// <param name="response">Callback action to handle coroutine string response</param>
        private IEnumerator GetCoroutine(string url, Action<string> response)
        {
            using(UnityWebRequest webRequest = UnityWebRequest.Get(url))
            {
                yield return webRequest.SendWebRequest();
                switch (webRequest.result)
                {
                    case UnityWebRequest.Result.ConnectionError:
                    case UnityWebRequest.Result.DataProcessingError:
                        break;
                    case UnityWebRequest.Result.ProtocolError:
                        Debug.LogError(url + ": HTTP Error: " + webRequest.error);
                        break;
                    case UnityWebRequest.Result.Success:
                        if (webRequest.GetResponseHeader("Content-Type") == "application/json")
                        {
                            response(webRequest.downloadHandler.text);
                        } else
                            Debug.LogError(url + " returned an invalid Content-Type of " + webRequest.GetResponseHeader("Content-Type"));
                        break;
                }
            }
        }

        /// <summary>
        ///     start a coroutine to Get image from an url.
        /// </summary>
        /// <param name="url">the url for the get request</param>
        /// <param name="response">Callback action to handle coroutine Texture2D response</param>
        public void GetTexture(string url, Action<Texture2D> response)
        {
            StartCoroutine(GetTextureCoroutine(url, response));
        }

        /// <summary>
        ///     IEnumerator used by GetTexture coroutine to execute the webrequest asynchronously
        /// </summary>
        /// <param name="url">the url for the get request</param>
        /// <param name="response">Callback action to handle coroutine Texture2D response</param>
        private IEnumerator GetTextureCoroutine(string url, Action<Texture2D> response)
        {
            using(UnityWebRequest webRequest = UnityWebRequestTexture.GetTexture(url))
            {
                yield return webRequest.SendWebRequest();
                switch (webRequest.result)
                {
                    case UnityWebRequest.Result.ConnectionError:
                    case UnityWebRequest.Result.DataProcessingError:
                        break;
                    case UnityWebRequest.Result.ProtocolError:
                        Debug.LogError(url + ": HTTP Error: " + webRequest.error);
                        break;
                    case UnityWebRequest.Result.Success:
                        response(((DownloadHandlerTexture)webRequest.downloadHandler).texture);
                        break;
                }
            }
        }
    }
}