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
                pub struct MessageEndpointOut {
                        /// List of message channels this endpoint listens to (omit for all).
                        #[serde(rename = "channels", skip_serializing_if = "Option::is_none")]
                        pub channels: Option<Vec<String>>,
                        #[serde(rename = "createdAt")]
                        pub created_at: String,
                        /// An example endpoint name.
                        #[serde(rename = "description")]
                        pub description: String,
                        #[serde(rename = "disabled", skip_serializing_if = "Option::is_none")]
                        pub disabled: Option<bool>,
                        #[serde(rename = "filterTypes", skip_serializing_if = "Option::is_none")]
                        pub filter_types: Option<Vec<String>>,
                        /// The ep's ID
                        #[serde(rename = "id")]
                        pub id: String,
                        #[serde(rename = "nextAttempt", skip_serializing_if = "Option::is_none")]
                        pub next_attempt: Option<String>,
                        #[serde(rename = "rateLimit", skip_serializing_if = "Option::is_none")]
                        pub rate_limit: Option<i32>,
                        #[serde(rename = "status")]
                        pub status: models::MessageStatus,
                        /// Optional unique identifier for the endpoint.
                        #[serde(rename = "uid", skip_serializing_if = "Option::is_none")]
                        pub uid: Option<String>,
                        #[serde(rename = "updatedAt")]
                        pub updated_at: String,
                        #[serde(rename = "url")]
                        pub url: String,
                        #[serde(rename = "version")]
                        pub version: i32,
                    }

                    impl MessageEndpointOut {
                    pub fn new(created_at: String, description: String, id: String, status: models::MessageStatus, updated_at: String, url: String, version: i32) -> MessageEndpointOut {
                MessageEndpointOut {
                    channels: None,
                    created_at,
                    description,
                    disabled: None,
                    filter_types: None,
                    id,
                    next_attempt: None,
                    rate_limit: None,
                    status,
                    uid: None,
                    updated_at,
                    url,
                    version,
                    }
                    }
                    }

