package com.svix.kotlin

import com.svix.kotlin.exceptions.ApiException
import com.svix.kotlin.internal.apis.MessageApi
import com.svix.kotlin.models.ApplicationIn
import com.svix.kotlin.models.ListResponseMessageOut
import com.svix.kotlin.models.MessageIn
import com.svix.kotlin.models.MessageOut
import java.time.OffsetDateTime

class MessageListOptions {
    var limit: Int? = null
    var iterator: String? = null
    var channel: String? = null
    var before: OffsetDateTime? = null
    var after: OffsetDateTime? = null
    var withContent: Boolean? = null
    var tag: String? = null
    var eventTypes: List<String>? = null

    /** Limit the number of returned items */
    fun limit(limit: Int) = apply { this.limit = limit }

    /** The iterator returned from a prior invocation */
    fun iterator(iterator: String) = apply { this.iterator = iterator }

    /** Filter response based on the channel. */
    fun channel(channel: String) = apply { this.channel = channel }

    /** Only include items created before a certain date. */
    fun before(before: OffsetDateTime) = apply { this.before = before }

    /** Only include items created after a certain date. */
    fun after(after: OffsetDateTime) = apply { this.after = after }

    /** When `true` message payloads are included in the response. */
    fun withContent(withContent: Boolean) = apply { this.withContent = withContent }

    /** Filter messages matching the provided tag. */
    fun tag(tag: String) = apply { this.tag = tag }

    /** Filter response based on the event type */
    fun eventTypes(eventTypes: List<String>) = apply { this.eventTypes = eventTypes }
}

class Message internal constructor(token: String, options: SvixOptions) {
    val api = MessageApi(options.serverUrl)

    init {
        api.accessToken = token
        api.userAgent = options.getUA()
        options.initialRetryDelayMillis?.let { api.initialRetryDelayMillis = it }
        options.numRetries?.let { api.numRetries = it }
    }

    suspend fun list(
        appId: String,
        options: MessageListOptions = MessageListOptions(),
    ): ListResponseMessageOut {
        try {
            return api.v1MessageList(
                appId,
                options.limit,
                options.iterator,
                options.channel,
                options.before,
                options.after,
                options.withContent,
                options.tag,
                HashSet(options.eventTypes),
            )
        } catch (e: Exception) {
            throw ApiException.wrap(e)
        }
    }

    suspend fun create(
        appId: String,
        messageIn: MessageIn,
        options: PostOptions = PostOptions(),
    ): MessageOut {
        try {
            return api.v1MessageCreate(appId, messageIn, null, options.idempotencyKey)
        } catch (e: Exception) {
            throw ApiException.wrap(e)
        }
    }

    suspend fun get(msgId: String, appId: String): MessageOut {
        try {
            return api.v1MessageGet(appId, msgId, null)
        } catch (e: Exception) {
            throw ApiException.wrap(e)
        }
    }

    suspend fun expungeContent(msgId: String, appId: String) {
        try {
            return api.v1MessageExpungeContent(appId, msgId)
        } catch (e: Exception) {
            throw ApiException.wrap(e)
        }
    }
}

/**
 * Creates a [MessageIn] with a pre-serialized payload.
 *
 * The payload is not normalized on the server. Normally, payloads are required to be JSON, and Svix
 * will minify the payload before sending the webhooks (for example, by removing extraneous
 * whitespace or unnecessarily escaped characters in strings). With this function, the payload will
 * be sent "as is", without any minification or other processing.
 *
 * @param payload Serialized message payload
 * @param contentType The value to use for the Content-Type header of the webhook sent by Svix,
 *   overwriting the default of `application/json` if specified
 *
 * See the class documentation for details about the other parameters.
 */
fun messageInRaw(
    eventType: String,
    payload: String,
    contentType: String? = null,
    application: ApplicationIn? = null,
    channels: Set<String>? = null,
    eventId: String? = null,
    payloadRetentionHours: Long? = null,
    payloadRetentionPeriod: Long? = 90L,
    tags: Set<String>? = null,
    transformationsParams: Map<String, Any> = mapOf(),
): MessageIn {
    val transformationsParams = transformationsParams.toMutableMap()
    transformationsParams.put("rawPayload", payload)
    if (contentType != null) {
        val headers = mapOf("content-type" to contentType)
        transformationsParams.put("headers", headers)
    }

    return MessageIn(
        eventType = eventType,
        payload = mapOf<String, String>(),
        application = application,
        channels = channels,
        eventId = eventId,
        payloadRetentionHours = payloadRetentionHours,
        payloadRetentionPeriod = payloadRetentionPeriod,
        tags = tags,
        transformationsParams = transformationsParams,
    )
}
