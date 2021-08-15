using System;
using System.Collections;
using UnityEngine;
using UnityEngine.Replay;
using UnityEngine.Networking;
using RestClient;

namespace Unity.Metacast.Demo
{
    /// <summary>
    ///     Populate UIBrowser with test json data
    /// </summary>
    public class TestListings : MonoBehaviour
    {
        // [SerializeField] private TextAsset m_TestJson;
        [SerializeField] private string baseUrl = "http://localhost:8080/api/v1/games";


        /// <summary>
        ///     Start is called on the frame when a script is enabled just
        ///     before any of the Update methods are called the first time.
        /// </summary>
        private void Start()
        {   
            HttpClient.Instance.Get(baseUrl, (string json) => {
                UIBrowser.instance.Init(json);
            });
        }
    }
}