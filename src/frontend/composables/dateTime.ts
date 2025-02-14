import dayjs, { type ManipulateType } from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';

const formatDateTime = (
  date: string | number,
  format: string = 'DD/MM/YYYY'
) => {
  if (!date) {
    return '';
  }

  return dayjs(date).format(format);
};

const unixDateTime = (date: string | number | Date | any) => {
  if (!date) {
    return '';
  }

  return dayjs(date).unix();
};

const readableDateTime = (date: string | number) => {
  if (!date) {
    return '';
  }

  dayjs.extend(relativeTime);

  return dayjs().to(dayjs(date));
};

const addDateTime = (
  date: string | number,
  count: number,
  unit: ManipulateType | undefined
) => {
  if (!date) {
    return '';
  }

  return dayjs(date).add(count, unit);
};

const isBefore = (date: string | number) => {
  if (!date) {
    return '';
  }

  const addedTime = addDateTime(date, 30, 'days');

  return dayjs().isBefore(addedTime, 'month');
};

export { formatDateTime, unixDateTime, readableDateTime, addDateTime, isBefore };
