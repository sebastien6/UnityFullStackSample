using System.Collections.Generic;

namespace UnityEngine.Replay
{
    /// <summary>
    ///     The top level browser for managing carousels and listings
    /// </summary>
    public class UIBrowser : MonoBehaviour
    {
        public static UIBrowser instance { get; private set; }

        /// <summary>
        ///  Awake is called when the script instance is being loaded.
        /// </summary>
        private void Awake()
        {
            instance = this;
        }

        /// <summary>
        /// Initialize the browser
        /// </summary>
        /// <param name="jsonString">The JSON string to parse into listings</param>
        public void Init(string jsonString)
        {
            CreateListings(jsonString);
        }

        /// <summary>
        ///     Creates the listings from a JSON string
        /// </summary>
        /// <param name="jsonString">The JSON string to parse</param>
        private void CreateListings(string jsonString)
        {
            ListingsContainer allListings = JsonUtility.FromJson<ListingsContainer>(jsonString);
            // List<Listing> listings = new List<Listing>();
            Dictionary<string, List<Listing>> byCat = new Dictionary<string, List<Listing>>();
            
            foreach (Listing listing in allListings.listings)
            {
                if (byCat.ContainsKey(listing.category)) {
                    byCat[listing.category].Add(listing);
                } else {
                    byCat.Add(listing.category, new List<Listing>());
                    byCat[listing.category].Add(listing);
                }
                
//                listings.Add(listing);
            }


            // foreach(var item in byCat){
                UICarousel carousel = FindObjectOfType<UICarousel>();
                if (carousel != null) {
                    // carousel.Init(item.Key, item.Value);
                    carousel.Init("live", byCat["live"]);
                } else {
                    Debug.LogError("UICarousel not found");
                }
            // }
        }
    }
}
