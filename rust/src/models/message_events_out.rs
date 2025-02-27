/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.1
 * 
 * Generated by: https://openapi-generator.tech
 */

#[allow(unused_imports)]
use crate::models;
#[allow(unused_imports)]
use serde::{Deserialize, Serialize};
        
                #[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
                pub struct MessageEventsOut {
                        #[serde(rename = "data")]
                        pub data: Vec<models::MessageOut>,
                        #[serde(rename = "done")]
                        pub done: bool,
                        #[serde(rename = "iterator")]
                        pub iterator: String,
                    }

                    impl MessageEventsOut {
                    pub fn new(data: Vec<models::MessageOut>, done: bool, iterator: String) -> MessageEventsOut {
                MessageEventsOut {
                    data,
                    done,
                    iterator,
                    }
                    }
                    }

