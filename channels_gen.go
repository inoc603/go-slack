package slack

// Auto-generated by internal/cmd/genmethods/genmethods.go. DO NOT EDIT!

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

var _ = strconv.Itoa
var _ = strings.Index
var _ = objects.EpochTime(0)

// ChannelsArchiveCall is created by ChannelsService.Archive method call
type ChannelsArchiveCall struct {
	service *ChannelsService
	channel string
}

// ChannelsCreateCall is created by ChannelsService.Create method call
type ChannelsCreateCall struct {
	service  *ChannelsService
	name     string
	validate bool
}

// ChannelsHistoryCall is created by ChannelsService.History method call
type ChannelsHistoryCall struct {
	service   *ChannelsService
	channel   string
	count     int // 1 - 1000
	inclusive bool
	latest    string // Range of time (end)
	oldest    string // Range of time (start)
	timestamp string // Used only when retrieving a single message
	unreads   bool   // Include unread_count_display in the output
}

// ChannelsInfoCall is created by ChannelsService.Info method call
type ChannelsInfoCall struct {
	service       *ChannelsService
	channel       string
	includeLocale bool
}

// ChannelsInviteCall is created by ChannelsService.Invite method call
type ChannelsInviteCall struct {
	service *ChannelsService
	channel string
	user    string
}

// ChannelsJoinCall is created by ChannelsService.Join method call
type ChannelsJoinCall struct {
	service  *ChannelsService
	name     string
	validate bool
}

// ChannelsKickCall is created by ChannelsService.Kick method call
type ChannelsKickCall struct {
	service *ChannelsService
	channel string
	user    string
}

// ChannelsLeaveCall is created by ChannelsService.Leave method call
type ChannelsLeaveCall struct {
	service *ChannelsService
	channel string
}

// ChannelsListCall is created by ChannelsService.List method call
type ChannelsListCall struct {
	service        *ChannelsService
	excludeArchive bool // Exclude archived channels
	excludeMembers bool // Exclude the list of members in channels
	limit          int
}

// ChannelsMarkCall is created by ChannelsService.Mark method call
type ChannelsMarkCall struct {
	service   *ChannelsService
	channel   string
	timestamp string
}

// ChannelsRenameCall is created by ChannelsService.Rename method call
type ChannelsRenameCall struct {
	service  *ChannelsService
	channel  string
	name     string
	validate bool
}

// ChannelsRepliesCall is created by ChannelsService.Replies method call
type ChannelsRepliesCall struct {
	service         *ChannelsService
	channel         string
	threadTimestamp string
}

// ChannelsSetPurposeCall is created by ChannelsService.SetPurpose method call
type ChannelsSetPurposeCall struct {
	service *ChannelsService
	channel string
	purpose string
}

// ChannelsSetTopicCall is created by ChannelsService.SetTopic method call
type ChannelsSetTopicCall struct {
	service *ChannelsService
	channel string
	topic   string
}

// ChannelsUnarchiveCall is created by ChannelsService.Unarchive method call
type ChannelsUnarchiveCall struct {
	service *ChannelsService
	channel string
}

// Archive creates a ChannelsArchiveCall object in preparation for accessing the channels.archive endpoint
func (s *ChannelsService) Archive(channel string) *ChannelsArchiveCall {
	var call ChannelsArchiveCall
	call.service = s
	call.channel = channel
	return &call
}

// ValidateArgs checks that all required fields are set in the ChannelsArchiveCall object
func (c *ChannelsArchiveCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	return nil
}

// Values returns the ChannelsArchiveCall object as url.Values
func (c *ChannelsArchiveCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("channel", c.channel)
	return v, nil
}

// Do executes the call to access channels.archive endpoint
func (c *ChannelsArchiveCall) Do(ctx context.Context) error {
	const endpoint = "channels.archive"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res struct {
		objects.GenericResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to channels.archive`)
	}
	if !res.OK {
		return errors.New(res.Error.String())
	}

	return nil
}

// FromValues parses the data in v and populates `c`
func (c *ChannelsArchiveCall) FromValues(v url.Values) error {
	var tmp ChannelsArchiveCall
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	*c = tmp
	return nil
}

// Create creates a ChannelsCreateCall object in preparation for accessing the channels.create endpoint
func (s *ChannelsService) Create(name string) *ChannelsCreateCall {
	var call ChannelsCreateCall
	call.service = s
	call.name = name
	return &call
}

// Validate sets the value for optional validate parameter
func (c *ChannelsCreateCall) Validate(validate bool) *ChannelsCreateCall {
	c.validate = validate
	return c
}

// ValidateArgs checks that all required fields are set in the ChannelsCreateCall object
func (c *ChannelsCreateCall) ValidateArgs() error {
	if len(c.name) <= 0 {
		return errors.New(`required field name not initialized`)
	}
	return nil
}

// Values returns the ChannelsCreateCall object as url.Values
func (c *ChannelsCreateCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("name", c.name)

	if c.validate {
		v.Set("validate", "true")
	}
	return v, nil
}

// Do executes the call to access channels.create endpoint
func (c *ChannelsCreateCall) Do(ctx context.Context) error {
	const endpoint = "channels.create"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res struct {
		objects.GenericResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to channels.create`)
	}
	if !res.OK {
		return errors.New(res.Error.String())
	}

	return nil
}

// FromValues parses the data in v and populates `c`
func (c *ChannelsCreateCall) FromValues(v url.Values) error {
	var tmp ChannelsCreateCall
	if raw := strings.TrimSpace(v.Get("name")); len(raw) > 0 {
		tmp.name = raw
	}
	if raw := strings.TrimSpace(v.Get("validate")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "validate"`)
		}
		tmp.validate = parsed
	}
	*c = tmp
	return nil
}

// History creates a ChannelsHistoryCall object in preparation for accessing the channels.history endpoint
func (s *ChannelsService) History(channel string) *ChannelsHistoryCall {
	var call ChannelsHistoryCall
	call.service = s
	call.channel = channel
	return &call
}

// Count sets the value for optional count parameter
func (c *ChannelsHistoryCall) Count(count int) *ChannelsHistoryCall {
	c.count = count
	return c
}

// Inclusive sets the value for optional inclusive parameter
func (c *ChannelsHistoryCall) Inclusive(inclusive bool) *ChannelsHistoryCall {
	c.inclusive = inclusive
	return c
}

// Latest sets the value for optional latest parameter
func (c *ChannelsHistoryCall) Latest(latest string) *ChannelsHistoryCall {
	c.latest = latest
	return c
}

// Oldest sets the value for optional oldest parameter
func (c *ChannelsHistoryCall) Oldest(oldest string) *ChannelsHistoryCall {
	c.oldest = oldest
	return c
}

// Timestamp sets the value for optional timestamp parameter
func (c *ChannelsHistoryCall) Timestamp(timestamp string) *ChannelsHistoryCall {
	c.timestamp = timestamp
	return c
}

// Unreads sets the value for optional unreads parameter
func (c *ChannelsHistoryCall) Unreads(unreads bool) *ChannelsHistoryCall {
	c.unreads = unreads
	return c
}

// ValidateArgs checks that all required fields are set in the ChannelsHistoryCall object
func (c *ChannelsHistoryCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	return nil
}

// Values returns the ChannelsHistoryCall object as url.Values
func (c *ChannelsHistoryCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("channel", c.channel)

	if c.count > 0 {
		v.Set("count", strconv.Itoa(c.count))
	}

	if c.inclusive {
		v.Set("inclusive", "true")
	}

	if len(c.latest) > 0 {
		v.Set("latest", c.latest)
	}

	if len(c.oldest) > 0 {
		v.Set("oldest", c.oldest)
	}

	if len(c.timestamp) > 0 {
		v.Set("ts", c.timestamp)
	}

	if c.unreads {
		v.Set("unreads", "true")
	}
	return v, nil
}

// Do executes the call to access channels.history endpoint
func (c *ChannelsHistoryCall) Do(ctx context.Context) (*objects.ChannelsHistoryResponse, error) {
	const endpoint = "channels.history"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.ChannelsHistoryResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to channels.history`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.ChannelsHistoryResponse, nil
}

// FromValues parses the data in v and populates `c`
func (c *ChannelsHistoryCall) FromValues(v url.Values) error {
	var tmp ChannelsHistoryCall
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	if raw := strings.TrimSpace(v.Get("count")); len(raw) > 0 {
		parsed, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return errors.Wrap(err, `failed to parse integer value "count"`)
		}
		tmp.count = int(parsed)
	}
	if raw := strings.TrimSpace(v.Get("inclusive")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "inclusive"`)
		}
		tmp.inclusive = parsed
	}
	if raw := strings.TrimSpace(v.Get("latest")); len(raw) > 0 {
		tmp.latest = raw
	}
	if raw := strings.TrimSpace(v.Get("oldest")); len(raw) > 0 {
		tmp.oldest = raw
	}
	if raw := strings.TrimSpace(v.Get("ts")); len(raw) > 0 {
		tmp.timestamp = raw
	}
	if raw := strings.TrimSpace(v.Get("unreads")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "unreads"`)
		}
		tmp.unreads = parsed
	}
	*c = tmp
	return nil
}

// Info creates a ChannelsInfoCall object in preparation for accessing the channels.info endpoint
func (s *ChannelsService) Info(channel string) *ChannelsInfoCall {
	var call ChannelsInfoCall
	call.service = s
	call.channel = channel
	return &call
}

// IncludeLocale sets the value for optional includeLocale parameter
func (c *ChannelsInfoCall) IncludeLocale(includeLocale bool) *ChannelsInfoCall {
	c.includeLocale = includeLocale
	return c
}

// ValidateArgs checks that all required fields are set in the ChannelsInfoCall object
func (c *ChannelsInfoCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	return nil
}

// Values returns the ChannelsInfoCall object as url.Values
func (c *ChannelsInfoCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("channel", c.channel)

	if c.includeLocale {
		v.Set("include_locale", "true")
	}
	return v, nil
}

// Do executes the call to access channels.info endpoint
func (c *ChannelsInfoCall) Do(ctx context.Context) (*objects.Channel, error) {
	const endpoint = "channels.info"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.Channel
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to channels.info`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.Channel, nil
}

// FromValues parses the data in v and populates `c`
func (c *ChannelsInfoCall) FromValues(v url.Values) error {
	var tmp ChannelsInfoCall
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	if raw := strings.TrimSpace(v.Get("include_locale")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "include_locale"`)
		}
		tmp.includeLocale = parsed
	}
	*c = tmp
	return nil
}

// Invite creates a ChannelsInviteCall object in preparation for accessing the channels.invite endpoint
func (s *ChannelsService) Invite(channel string, user string) *ChannelsInviteCall {
	var call ChannelsInviteCall
	call.service = s
	call.channel = channel
	call.user = user
	return &call
}

// ValidateArgs checks that all required fields are set in the ChannelsInviteCall object
func (c *ChannelsInviteCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	if len(c.user) <= 0 {
		return errors.New(`required field user not initialized`)
	}
	return nil
}

// Values returns the ChannelsInviteCall object as url.Values
func (c *ChannelsInviteCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("channel", c.channel)

	v.Set("user", c.user)
	return v, nil
}

// Do executes the call to access channels.invite endpoint
func (c *ChannelsInviteCall) Do(ctx context.Context) (*objects.Channel, error) {
	const endpoint = "channels.invite"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.Channel
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to channels.invite`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.Channel, nil
}

// FromValues parses the data in v and populates `c`
func (c *ChannelsInviteCall) FromValues(v url.Values) error {
	var tmp ChannelsInviteCall
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	if raw := strings.TrimSpace(v.Get("user")); len(raw) > 0 {
		tmp.user = raw
	}
	*c = tmp
	return nil
}

// Join creates a ChannelsJoinCall object in preparation for accessing the channels.join endpoint
func (s *ChannelsService) Join(name string) *ChannelsJoinCall {
	var call ChannelsJoinCall
	call.service = s
	call.name = name
	return &call
}

// Validate sets the value for optional validate parameter
func (c *ChannelsJoinCall) Validate(validate bool) *ChannelsJoinCall {
	c.validate = validate
	return c
}

// ValidateArgs checks that all required fields are set in the ChannelsJoinCall object
func (c *ChannelsJoinCall) ValidateArgs() error {
	if len(c.name) <= 0 {
		return errors.New(`required field name not initialized`)
	}
	return nil
}

// Values returns the ChannelsJoinCall object as url.Values
func (c *ChannelsJoinCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("name", c.name)

	if c.validate {
		v.Set("validate", "true")
	}
	return v, nil
}

// Do executes the call to access channels.join endpoint
func (c *ChannelsJoinCall) Do(ctx context.Context) (*objects.Channel, error) {
	const endpoint = "channels.join"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.Channel
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to channels.join`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.Channel, nil
}

// FromValues parses the data in v and populates `c`
func (c *ChannelsJoinCall) FromValues(v url.Values) error {
	var tmp ChannelsJoinCall
	if raw := strings.TrimSpace(v.Get("name")); len(raw) > 0 {
		tmp.name = raw
	}
	if raw := strings.TrimSpace(v.Get("validate")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "validate"`)
		}
		tmp.validate = parsed
	}
	*c = tmp
	return nil
}

// Kick creates a ChannelsKickCall object in preparation for accessing the channels.kick endpoint
func (s *ChannelsService) Kick(channel string, user string) *ChannelsKickCall {
	var call ChannelsKickCall
	call.service = s
	call.channel = channel
	call.user = user
	return &call
}

// ValidateArgs checks that all required fields are set in the ChannelsKickCall object
func (c *ChannelsKickCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	if len(c.user) <= 0 {
		return errors.New(`required field user not initialized`)
	}
	return nil
}

// Values returns the ChannelsKickCall object as url.Values
func (c *ChannelsKickCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("channel", c.channel)

	v.Set("user", c.user)
	return v, nil
}

// Do executes the call to access channels.kick endpoint
func (c *ChannelsKickCall) Do(ctx context.Context) error {
	const endpoint = "channels.kick"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res struct {
		objects.GenericResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to channels.kick`)
	}
	if !res.OK {
		return errors.New(res.Error.String())
	}

	return nil
}

// FromValues parses the data in v and populates `c`
func (c *ChannelsKickCall) FromValues(v url.Values) error {
	var tmp ChannelsKickCall
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	if raw := strings.TrimSpace(v.Get("user")); len(raw) > 0 {
		tmp.user = raw
	}
	*c = tmp
	return nil
}

// Leave creates a ChannelsLeaveCall object in preparation for accessing the channels.leave endpoint
func (s *ChannelsService) Leave(channel string) *ChannelsLeaveCall {
	var call ChannelsLeaveCall
	call.service = s
	call.channel = channel
	return &call
}

// ValidateArgs checks that all required fields are set in the ChannelsLeaveCall object
func (c *ChannelsLeaveCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	return nil
}

// Values returns the ChannelsLeaveCall object as url.Values
func (c *ChannelsLeaveCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("channel", c.channel)
	return v, nil
}

// Do executes the call to access channels.leave endpoint
func (c *ChannelsLeaveCall) Do(ctx context.Context) error {
	const endpoint = "channels.leave"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res struct {
		objects.GenericResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to channels.leave`)
	}
	if !res.OK {
		return errors.New(res.Error.String())
	}

	return nil
}

// FromValues parses the data in v and populates `c`
func (c *ChannelsLeaveCall) FromValues(v url.Values) error {
	var tmp ChannelsLeaveCall
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	*c = tmp
	return nil
}

// List creates a ChannelsListCall object in preparation for accessing the channels.list endpoint
func (s *ChannelsService) List() *ChannelsListCall {
	var call ChannelsListCall
	call.service = s
	return &call
}

// ExcludeArchive sets the value for optional excludeArchive parameter
func (c *ChannelsListCall) ExcludeArchive(excludeArchive bool) *ChannelsListCall {
	c.excludeArchive = excludeArchive
	return c
}

// ExcludeMembers sets the value for optional excludeMembers parameter
func (c *ChannelsListCall) ExcludeMembers(excludeMembers bool) *ChannelsListCall {
	c.excludeMembers = excludeMembers
	return c
}

// Limit sets the value for optional limit parameter
func (c *ChannelsListCall) Limit(limit int) *ChannelsListCall {
	c.limit = limit
	return c
}

// ValidateArgs checks that all required fields are set in the ChannelsListCall object
func (c *ChannelsListCall) ValidateArgs() error {
	return nil
}

// Values returns the ChannelsListCall object as url.Values
func (c *ChannelsListCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.excludeArchive {
		v.Set("exclude_archive", "true")
	}

	if c.excludeMembers {
		v.Set("exclude_members", "true")
	}

	if c.limit > 0 {
		v.Set("limit", strconv.Itoa(c.limit))
	}
	return v, nil
}

// Do executes the call to access channels.list endpoint
func (c *ChannelsListCall) Do(ctx context.Context) (objects.ChannelList, error) {
	const endpoint = "channels.list"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		objects.ChannelList `json:"channels"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to channels.list`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.ChannelList, nil
}

// FromValues parses the data in v and populates `c`
func (c *ChannelsListCall) FromValues(v url.Values) error {
	var tmp ChannelsListCall
	if raw := strings.TrimSpace(v.Get("exclude_archive")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "exclude_archive"`)
		}
		tmp.excludeArchive = parsed
	}
	if raw := strings.TrimSpace(v.Get("exclude_members")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "exclude_members"`)
		}
		tmp.excludeMembers = parsed
	}
	if raw := strings.TrimSpace(v.Get("limit")); len(raw) > 0 {
		parsed, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return errors.Wrap(err, `failed to parse integer value "limit"`)
		}
		tmp.limit = int(parsed)
	}
	*c = tmp
	return nil
}

// Mark creates a ChannelsMarkCall object in preparation for accessing the channels.mark endpoint
func (s *ChannelsService) Mark(channel string) *ChannelsMarkCall {
	var call ChannelsMarkCall
	call.service = s
	call.channel = channel
	return &call
}

// Timestamp sets the value for optional timestamp parameter
func (c *ChannelsMarkCall) Timestamp(timestamp string) *ChannelsMarkCall {
	c.timestamp = timestamp
	return c
}

// ValidateArgs checks that all required fields are set in the ChannelsMarkCall object
func (c *ChannelsMarkCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	return nil
}

// Values returns the ChannelsMarkCall object as url.Values
func (c *ChannelsMarkCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("channel", c.channel)

	if len(c.timestamp) > 0 {
		v.Set("ts", c.timestamp)
	}
	return v, nil
}

// Do executes the call to access channels.mark endpoint
func (c *ChannelsMarkCall) Do(ctx context.Context) error {
	const endpoint = "channels.mark"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res struct {
		objects.GenericResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to channels.mark`)
	}
	if !res.OK {
		return errors.New(res.Error.String())
	}

	return nil
}

// FromValues parses the data in v and populates `c`
func (c *ChannelsMarkCall) FromValues(v url.Values) error {
	var tmp ChannelsMarkCall
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	if raw := strings.TrimSpace(v.Get("ts")); len(raw) > 0 {
		tmp.timestamp = raw
	}
	*c = tmp
	return nil
}

// Rename creates a ChannelsRenameCall object in preparation for accessing the channels.rename endpoint
func (s *ChannelsService) Rename(channel string, name string) *ChannelsRenameCall {
	var call ChannelsRenameCall
	call.service = s
	call.channel = channel
	call.name = name
	return &call
}

// Validate sets the value for optional validate parameter
func (c *ChannelsRenameCall) Validate(validate bool) *ChannelsRenameCall {
	c.validate = validate
	return c
}

// ValidateArgs checks that all required fields are set in the ChannelsRenameCall object
func (c *ChannelsRenameCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	if len(c.name) <= 0 {
		return errors.New(`required field name not initialized`)
	}
	return nil
}

// Values returns the ChannelsRenameCall object as url.Values
func (c *ChannelsRenameCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("channel", c.channel)

	v.Set("name", c.name)

	if c.validate {
		v.Set("validate", "true")
	}
	return v, nil
}

// Do executes the call to access channels.rename endpoint
func (c *ChannelsRenameCall) Do(ctx context.Context) (*objects.Channel, error) {
	const endpoint = "channels.rename"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.Channel `json:"channel"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to channels.rename`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.Channel, nil
}

// FromValues parses the data in v and populates `c`
func (c *ChannelsRenameCall) FromValues(v url.Values) error {
	var tmp ChannelsRenameCall
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	if raw := strings.TrimSpace(v.Get("name")); len(raw) > 0 {
		tmp.name = raw
	}
	if raw := strings.TrimSpace(v.Get("validate")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "validate"`)
		}
		tmp.validate = parsed
	}
	*c = tmp
	return nil
}

// Replies creates a ChannelsRepliesCall object in preparation for accessing the channels.replies endpoint
func (s *ChannelsService) Replies(channel string, threadTimestamp string) *ChannelsRepliesCall {
	var call ChannelsRepliesCall
	call.service = s
	call.channel = channel
	call.threadTimestamp = threadTimestamp
	return &call
}

// ValidateArgs checks that all required fields are set in the ChannelsRepliesCall object
func (c *ChannelsRepliesCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	if len(c.threadTimestamp) <= 0 {
		return errors.New(`required field threadTimestamp not initialized`)
	}
	return nil
}

// Values returns the ChannelsRepliesCall object as url.Values
func (c *ChannelsRepliesCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("channel", c.channel)

	v.Set("thread_ts", c.threadTimestamp)
	return v, nil
}

// Do executes the call to access channels.replies endpoint
func (c *ChannelsRepliesCall) Do(ctx context.Context) (objects.MessageList, error) {
	const endpoint = "channels.replies"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		objects.MessageList `json:"messages"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to channels.replies`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.MessageList, nil
}

// FromValues parses the data in v and populates `c`
func (c *ChannelsRepliesCall) FromValues(v url.Values) error {
	var tmp ChannelsRepliesCall
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	if raw := strings.TrimSpace(v.Get("thread_ts")); len(raw) > 0 {
		tmp.threadTimestamp = raw
	}
	*c = tmp
	return nil
}

// SetPurpose creates a ChannelsSetPurposeCall object in preparation for accessing the channels.setPurpose endpoint
func (s *ChannelsService) SetPurpose(channel string, purpose string) *ChannelsSetPurposeCall {
	var call ChannelsSetPurposeCall
	call.service = s
	call.channel = channel
	call.purpose = purpose
	return &call
}

// ValidateArgs checks that all required fields are set in the ChannelsSetPurposeCall object
func (c *ChannelsSetPurposeCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	if len(c.purpose) <= 0 {
		return errors.New(`required field purpose not initialized`)
	}
	return nil
}

// Values returns the ChannelsSetPurposeCall object as url.Values
func (c *ChannelsSetPurposeCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("channel", c.channel)

	v.Set("purpose", c.purpose)
	return v, nil
}

// Do executes the call to access channels.setPurpose endpoint
func (c *ChannelsSetPurposeCall) Do(ctx context.Context) (*string, error) {
	const endpoint = "channels.setPurpose"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*string `json:"purpose"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to channels.setPurpose`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.string, nil
}

// FromValues parses the data in v and populates `c`
func (c *ChannelsSetPurposeCall) FromValues(v url.Values) error {
	var tmp ChannelsSetPurposeCall
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	if raw := strings.TrimSpace(v.Get("purpose")); len(raw) > 0 {
		tmp.purpose = raw
	}
	*c = tmp
	return nil
}

// SetTopic creates a ChannelsSetTopicCall object in preparation for accessing the channels.setTopic endpoint
func (s *ChannelsService) SetTopic(channel string, topic string) *ChannelsSetTopicCall {
	var call ChannelsSetTopicCall
	call.service = s
	call.channel = channel
	call.topic = topic
	return &call
}

// ValidateArgs checks that all required fields are set in the ChannelsSetTopicCall object
func (c *ChannelsSetTopicCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	if len(c.topic) <= 0 {
		return errors.New(`required field topic not initialized`)
	}
	return nil
}

// Values returns the ChannelsSetTopicCall object as url.Values
func (c *ChannelsSetTopicCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("channel", c.channel)

	v.Set("topic", c.topic)
	return v, nil
}

// Do executes the call to access channels.setTopic endpoint
func (c *ChannelsSetTopicCall) Do(ctx context.Context) (*string, error) {
	const endpoint = "channels.setTopic"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*string `json:"topic"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to channels.setTopic`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.string, nil
}

// FromValues parses the data in v and populates `c`
func (c *ChannelsSetTopicCall) FromValues(v url.Values) error {
	var tmp ChannelsSetTopicCall
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	if raw := strings.TrimSpace(v.Get("topic")); len(raw) > 0 {
		tmp.topic = raw
	}
	*c = tmp
	return nil
}

// Unarchive creates a ChannelsUnarchiveCall object in preparation for accessing the channels.unarchive endpoint
func (s *ChannelsService) Unarchive(channel string) *ChannelsUnarchiveCall {
	var call ChannelsUnarchiveCall
	call.service = s
	call.channel = channel
	return &call
}

// ValidateArgs checks that all required fields are set in the ChannelsUnarchiveCall object
func (c *ChannelsUnarchiveCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	return nil
}

// Values returns the ChannelsUnarchiveCall object as url.Values
func (c *ChannelsUnarchiveCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("channel", c.channel)
	return v, nil
}

// Do executes the call to access channels.unarchive endpoint
func (c *ChannelsUnarchiveCall) Do(ctx context.Context) error {
	const endpoint = "channels.unarchive"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res struct {
		objects.GenericResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to channels.unarchive`)
	}
	if !res.OK {
		return errors.New(res.Error.String())
	}

	return nil
}

// FromValues parses the data in v and populates `c`
func (c *ChannelsUnarchiveCall) FromValues(v url.Values) error {
	var tmp ChannelsUnarchiveCall
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	*c = tmp
	return nil
}
