**progress**
66/69
```
Y   #define 	GST_EVENT_MAKE_TYPE()
Y   #define 	GST_EVENT_TYPE()
Y   #define 	GST_EVENT_TYPE_NAME()
Y   #define 	GST_EVENT_TIMESTAMP()
Y   #define 	GST_EVENT_SEQNUM()
Y   #define 	GST_EVENT_IS_UPSTREAM()
Y   #define 	GST_EVENT_IS_DOWNSTREAM()
Y   #define 	GST_EVENT_IS_SERIALIZED()
Y   #define 	GST_EVENT_IS_STICKY()
Y   GstEventTypeFlags 	gst_event_type_get_flags ()
Y   const gchar * 	gst_event_type_get_name ()
Y   GQuark 	gst_event_type_to_quark ()
Y   GstEvent * 	gst_event_ref ()
Y   void 	gst_event_unref ()
    gboolean 	gst_event_replace ()
Y   GstEvent * 	gst_event_copy ()
    GstEvent * 	gst_event_steal ()
    gboolean 	gst_event_take ()
Y   #define 	gst_event_is_writable()
Y   #define 	gst_event_make_writable()
Y   GstStructure * 	gst_event_writable_structure ()
Y   GstEvent * 	gst_event_new_custom ()
Y   const GstStructure * 	gst_event_get_structure ()
Y   gboolean 	gst_event_has_name ()
Y   guint32 	gst_event_get_seqnum ()
Y   void 	gst_event_set_seqnum ()
Y   gint64 	gst_event_get_running_time_offset ()
Y   void 	gst_event_set_running_time_offset ()
Y   GstEvent * 	gst_event_new_flush_start ()
Y   GstEvent * 	gst_event_new_flush_stop ()
Y   void 	gst_event_parse_flush_stop ()
Y   GstEvent * 	gst_event_new_eos ()
Y   GstEvent * 	gst_event_new_gap ()
Y   void 	gst_event_parse_gap ()
Y   GstEvent * 	gst_event_new_stream_start ()
Y   void 	gst_event_parse_stream_start ()
Y   void 	gst_event_set_stream_flags ()
Y   void 	gst_event_parse_stream_flags ()
Y   void 	gst_event_set_group_id ()
Y   gboolean 	gst_event_parse_group_id ()
Y   GstEvent * 	gst_event_new_segment ()
Y   void 	gst_event_parse_segment ()
Y   void 	gst_event_copy_segment ()
Y   GstEvent * 	gst_event_new_tag ()
Y   void 	gst_event_parse_tag ()
Y   GstEvent * 	gst_event_new_buffer_size ()
Y   void 	gst_event_parse_buffer_size ()
Y   GstEvent * 	gst_event_new_qos ()
Y   void 	gst_event_parse_qos ()
Y   GstEvent * 	gst_event_new_seek ()
Y   void 	gst_event_parse_seek ()
Y   GstEvent * 	gst_event_new_navigation ()
Y   GstEvent * 	gst_event_new_latency ()
Y   void 	gst_event_parse_latency ()
Y   GstEvent * 	gst_event_new_step ()
Y   void 	gst_event_parse_step ()
Y   GstEvent * 	gst_event_new_sink_message ()
Y   void 	gst_event_parse_sink_message ()
Y   GstEvent * 	gst_event_new_reconfigure ()
Y   GstEvent * 	gst_event_new_caps ()
Y   void 	gst_event_parse_caps ()
Y   GstEvent * 	gst_event_new_toc ()
Y   void 	gst_event_parse_toc ()
Y   GstEvent * 	gst_event_new_toc_select ()
Y   void 	gst_event_parse_toc_select ()
Y   GstEvent * 	gst_event_new_segment_done ()
Y   void 	gst_event_parse_segment_done ()
Y   GstEvent * 	gst_event_new_protection ()
Y   void 	gst_event_parse_protection ()
```